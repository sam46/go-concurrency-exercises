//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import "time"

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

func notifyDone(f func()) <-chan struct{} {
	c := make(chan struct{})
	go func(c chan struct{}) {
		f()
		close(c)
	}(c)
	return c
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	done := notifyDone(process)
	timeout := time.After(time.Second * 10)
	select {
	case <-done:
		return true
	case <-timeout:
		break
	}
	if !u.IsPremium {
		// call process.kill() or similar to stop it here.
		// but right now the process, as given, can't be stopped
		// because it doesn't take a context, or provide some mechanism for stopping it.
		// (assuming that modifying the given params' type/shape/etc isn't allowed)
		return false
	}
	return true
}

func main() {
	RunMockServer()
}
