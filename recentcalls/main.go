package main

func main() {
	// ["RecentCounter", "ping", "ping", "ping", "ping"]
	// [[], [1], [100], [3001], [3002]]
	// Output: [null, 1, 2, 3, 3]
	obj := Constructor()
	println(obj.Ping(1))    // Output: 1
	println(obj.Ping(100))  // Output: 2
	println(obj.Ping(3001)) // Output: 3
	println(obj.Ping(3002)) // Output: 3
}

type RecentCounter struct {
	calls []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		calls: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.calls = append(this.calls, t)
	for i := len(this.calls) - 2; i >= 0; i-- {
		if t-this.calls[i] > 3000 {
			this.calls = this.calls[i+1:]
			break
		}
	}
	return len(this.calls)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

// Notes:
// Redis would be interesting for a real-world application, but not for a story problem scenario.
// Focus on in-memory solutions.
//
// "It is guaranteed that every call to ping uses a strictly larger value of t than the previous call."
// This is telling us two important things:
// - Each value of t is unique
// - Any t values < t-3000 will never be needed again

// ping method -> returns a count of requests from the lastt 3000 ms
// so: we need to keep an ordered slice of times
// on each call to ping:
// - start on the right side of the slice
// - move left until end is reached, or a value is found that's < t-3000
// discard any extra values on the left
// return count of elements

// complexity: O(n) - Where n is the number of invocations within the 3000ms range.
// main problem that tripped me up: not reading instructions carefully.
// I needed to make the result inclusive of the current ping.
