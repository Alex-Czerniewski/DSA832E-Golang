package main

import ("fmt"; "net"; "time")

func main() {
    ip := "192.168.0.108"
    port := "5555"
    ip_port := ip+":"+port

    conn, err := net.Dial("tcp", fmt.Sprintf(ip_port))
    if err != nil {
        fmt.Println("Connection error:", err)
        return
    }
    defer conn.Close()

    conn.SetReadDeadline(time.Now().Add(5 * time.Second))

    idn := "*IDN?\n"
    _, err = conn.Write([]byte(idn))
    if err != nil {
        fmt.Println("Write error:", err)
        return
    }

    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil {
        fmt.Println("Read error:", err)
        return
    }

    fmt.Printf("%s\n", buf[:n])

    Centercmd := "SENS:FREQ:CENT 1000000\n"
    Spancmd := "SENS:FREQ:SPAN 10000000\n"

    _, err = conn.Write([]byte(Centercmd))
    if err != nil {
        fmt.Println("Write error:", err)
        return
    }

    _, err = conn.Write([]byte(Spancmd))
    if err != nil {
        fmt.Println("Write error:", err)
        return
    }

    F := 1000000

    for i := 0; i < 10; i++ {
        F += 100000000
        cmd := fmt.Sprintf("SENS:FREQ:CENT %d\n", F)
        fmt.Println(cmd)
        _, err = conn.Write([]byte(cmd))
        if err != nil {
            fmt.Println("Error sending command:", err)
            return
        }
        time.Sleep(5 * time.Second)
    }
}
