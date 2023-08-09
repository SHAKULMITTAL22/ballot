func TestRunTest(t *testing.T) {
    var (
        w       http.ResponseWriter
        r       *http.Request
        status  Status
        err     error
    )

    // Test successful execution
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    defer r.Body.Close()
    log.Println("ballot endpoint tests running")
    status.Message = "Test Cases passed"
    status.Code = http.StatusOK
    writeVoterResponse(w, status)

    // Test failed execution with error
    err = errors.New("test error")
    status.Message = fmt.Sprint("Test Cases Failed with error : ", err)
    status.Code = http.StatusBadRequest
    writeVoterResponse(w, status)
}
