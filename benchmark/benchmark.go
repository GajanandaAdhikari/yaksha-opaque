startTotal := time.Now()

// Loading Phase
startLoad := time.Now()
// Load protocol parameters, keys etc.
loadProtocolParams()
elapsedLoad := time.Since(startLoad)

// Key Exchange Phase
startCalc := time.Now()
// Perform local cryptographic calculation (e.g., OPRF blind)
clientCalc()
elapsedCalc := time.Since(startCalc)

// Send request and wait for response from server
startLatency := time.Now()
sendRequest()
receiveResponse()
elapsedLatency := time.Since(startLatency)

totalElapsed := time.Since(startTotal)

fmt.Printf("Loading Time: %v\n", elapsedLoad)
fmt.Printf("Calculation Time: %v\n", elapsedCalc)
fmt.Printf("Latency Time: %v\n", elapsedLatency)
fmt.Printf("Total Benchmark Time: %v\n", totalElapsed)
