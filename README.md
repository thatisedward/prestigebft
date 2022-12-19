<h1 align="center"> Vintage BFT: Making View Changes Efficient for Byzantine Fault-Tolerant
Consensus Algorithms using Reputation Mechanisms  </h1>

## About a typo!
We are sorry to clarify a typo in the paper that causes misunderstanding.

On page 5, the second last paragraph, "... the reputation engine does incur ..."; this should be "the reputation engine does **not** incur ..."

## Dependencies
This project uses two external packages for threshold signatures and logging.

    # threshold signatures
    go get github.com/dedis/kyber/
    # logging
    go get github.com/sirupsen/logrus


## Important parameters

`-b int` is the batch size\
`-mq int` is the maximum size of message queues\
`-p int` is the number of packing threads used in one consensus instance\
`-pt int` is the number of computing threads used in view changes\

`-n int` is the number of servers in total (system configuration)\
`-th int` is the threshold value for threshold signatures (quorums)\
`-id int` is this server's ID\
`-d int` is the emulated network delay\

`-ns bool` enables native storage for storing committed entries in plain text files\
`-r bool` enables log responsiveness\
`-gc bool` enables deleting cache entries that have been committed by consensus\
`-repu bool` enables the reputation engine with computing hash computaional work\

`-s int` inital server state: 0 : Leader;  1 : Nominee; 2 : Starlet; 3 : Worker\
`-rp bool` prints real time log on the screen\
`perf` enables peak prformance testing configuration, which disables `rp` and `gc`\

## Deployment
The project is under a double-blind review process. To prevent our deployment details from disclosing our institution information, deployment scripts are temporarily redacted to preserve anonymity.
