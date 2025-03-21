package daemon

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilwait "k8s.io/apimachinery/pkg/util/wait"
)

type ReadyTracker struct {
	mutex          sync.Mutex
	config         bool
	processManager *ProcessManager
}

func (rt *ReadyTracker) Ready() (bool, string) {
	rt.mutex.Lock()
	defer rt.mutex.Unlock()

	if !rt.config {
		return false, "Config not applied"
	}

	if len(rt.processManager.process) == 0 {
		return false, "No processes have started"
	}

	notRunning := strings.Builder{}
	noMetrics := strings.Builder{}
	for _, p := range rt.processManager.process {
		if p.Stopped() {
			if notRunning.Len() > 0 {
				notRunning.WriteString(", ")
			}
			notRunning.WriteString(p.name)
		} else if !p.hasCollectedMetrics {
			if noMetrics.Len() > 0 {
				noMetrics.WriteString(", ")
			}
			noMetrics.WriteString(p.name)
		}

	}
	if notRunning.Len() > 0 {
		return false, "Stopped process(es): " + notRunning.String()
	}

	if noMetrics.Len() > 0 {
		return false, "Process(es) have not yet collected metrics: " + noMetrics.String()
	}

	return true, ""
}

func (rt *ReadyTracker) setConfig(v bool) {
	rt.mutex.Lock()
	rt.config = v
	rt.mutex.Unlock()
}

type readyHandler struct {
	tracker *ReadyTracker
}

func (h readyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if isReady, msg := h.tracker.Ready(); !isReady {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "503: %s\n", msg)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func StartReadyServer(bindAddress string, tracker *ReadyTracker) {
	glog.Info("Starting Ready Server")
	mux := http.NewServeMux()
	mux.Handle("/ready", readyHandler{tracker: tracker})
	go utilwait.Until(func() {
		err := http.ListenAndServe(bindAddress, mux)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("starting metrics server failed: %v", err))
		}
	}, 5*time.Second, utilwait.NeverStop)
}
