package handlers

import (
	"net/url"
	"sync"
)

// State is the discovery server configuration
// state shared between handlers.
type State struct {
	mu            sync.RWMutex
	etcdHost      string
	etcdCURL      *url.URL
	currentLeader string
	discHost      string
}

func (st *State) endpoint() (ep string) {
	st.mu.RLock()
	ep = st.etcdHost
	st.mu.RUnlock()
	return ep
}

func (st *State) getCurrentLeader() (leader string) {
	st.mu.RLock()
	leader = st.currentLeader
	st.mu.RUnlock()
	return leader
}

func (st *State) setCurrentLeader(leader string) {
	st.mu.Lock()
	st.currentLeader = leader
	st.mu.Unlock()
}
