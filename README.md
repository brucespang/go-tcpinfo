# go-tcpinfo

This is a small wrapper around the syscall library, so you don't have to mess around with it whenever you want to get tcp info for a connection

## Usage

```go
tcpInfo, err := tcpinfo.GetsockoptTCPInfo(&conn)
if err != nil {
    panic(err)
}
```

## Example

```
$ go run example.go
&{State:1 Ca_state:0 Retransmits:0 Probes:0 Backoff:0 Options:7 Pad_cgo_0:[153 0] Rto:204000 Ato:0 Snd_mss:22016 Rcv_mss:536 Unacked:0 Sacked:0 Lost:0 Retrans:0 Fackets:0 Last_data_sent:0 Last_ack_sent:0 Last_data_recv:3828100 Last_ack_recv:0 Pmtu:65535 Rcv_ssthresh:43690 Rtt:4000 Rttvar:2000 Snd_ssthresh:2147483647 Snd_cwnd:10 Advmss:65483 Reordering:3 Rcv_rtt:0 Rcv_space:43690 Total_retrans:0}
```
