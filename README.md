<h1 align="center"> Prestige BFT: Making Decentralization Efficient in Distributed Ledgers using Reputation-based Byzantine Fault-Tolerant Consensus Algorithms  </h1>

## About Sigmod submission!

We are sorry to first darw reviewers attention to clarify a typo in the submission that may lead to misunderstanding.

On page 7, the last sentence of the first paragraph (S_a <3>): "Thus, S_c receives no penalty ..." This should be `S_a`, there is no S_c.

## About PrestigeBFT

PrestigeBFT is a reputation-based BFT consensus algorithm that enables active view changes in BFT algorithms. Under the BFT model, we cannot pass an absolute judgment on server correctness (e.g., we never know if a server deliberately drops a message or if the message is dropped by the network). Instead, **PrestigeBFT extends the traditional state machine replication (SMR) properties and establishes a reputation state to describe a server's likelihood of being correct or faulty.** Furthermore, PrestigeBFT develops a reputation engine that precisely translates server behavior history into a reputation penalty. A server's reputation penalty determines the difficulty of the server becoming a future leader via a Proof-of-Work-like computation phase. As such, more correct servers are more likely to become future leaders as they perform less computation compared to faulty servers. Equipped with the reputation engine, **PrestigeBFT enables active view changes where servers can proactively campaign for leadership for themselves.** Compared with PBFT-like BFT algorithms (e.g., Zyzzyva, SBFT, and HotStuff), 1) under benign failures (e.g., crash failures), PrestigeBFT does not blindly follow a pre-defined leadership schedule and thus prevents the system from having crashed servers as a leader; 2) under Byzantine failures, PrestigeBFT gradually imposes increasing computational work on faulty servers and thus suppresses Byzantine servers from becoming leaders over time.

## Q & A Prestige BFT
We've collected some questions about Prestige BFT during from presentations, lectures, and conversions with ECE/CS students, professors, and DS developers. We aim to help in understanding the key concepts of Prestige BFT by showing these questions.

1. **[Fault Tolerance]** *The paper discussed that it is not a good approach to kick out a server whenever it fails in Section 1.2. Could you further elaborate on this?*

*Answer:* We'd like to discuss this question from two perspectives. **First, kicking out a server is a system reconfiguration approach, which is no longer fault tolerance anymore.** In consensus algorithms, the key idea is to tolerate failures, which means the system can operate correctly while some servers exhibit failures. In BFT, the system can still operate correctly with some servers being compromised. Fault tolerance is superior to reconfiguration because it keeps the system intact.\
\
Second, under partial synchrony, **kicking out a faulty server requires constant and manual reconfigurations, and the system may quickly run out of servers.** A server can exhibit failures simply because it encounters a network failure (benign), such as packet loss or duplication. In this case, if this server is kicked out because of a benign failure, then in a large-scale system, this scenario may become common and operators will have to constantly reconfigure the system. According to the FLP theorem, we never know if a server deliberately abandons a request or if the request is delayed/dropped by the network. Thus, consensus algorithms rely on fault tolerance, and Prestige BFT strengthens system robustness by improving fault tolerance.

2. **[Consensus Protocols]** *Will the computation calculation affect transaction replication?*

*Answer:* **The short answer is NO. The computation calculation never affects transaction replication.** A server operates under the replication protocol in the worker state and only performs computation when it has transitioned to the starlet state. Replication and performing computation never coexist. A worker transitions to a starlet when it detects a leader's failure; at this point, the replication has already stopped (because the leader has failed). Additionally, as shown in `Figure 3`, a starlet will go back to a worker whenever it finds a legitimate leader (e.g., some other server is elected as a new leader). It aborts its computation and operates the replication protocol under the new leadership.

3. **[Reputation Engine]** *Can the reputation engine calculate server reputation based on different criteria?*

*Answer:* **The short answer is YES. Prestige BFT proposes a general BFT consensus architecture that extends the traditional BFT SMR properties to a reputation state.** This reputation state can be implemented differently as long as it is indicative of server correctness, i.e., being malicious or being correct. Our reputation engine is one way to translate a server's past behavior into a reputation penalty that determines how difficult this server can become the next leader.

4. **[Reputation Engine]** *What if faulty servers have much stronger computational power than correct servers? How will this affect Prestige BFT?*

*Answer:* **Of course a more powerful faulty server or a cohort of faulty servers (f>1) can cause more "damage", but they cannot avoid the exponential computation cost.** As we have shown in the evaluation under Byzantine failures that when the system has more than one faulty server, Byzantine servers can collude to perform computation and amortize the penalization, imposed by the reputation engine, among them. Thus, a more powerful faulty server (or a cohort of faulty servers) may perform a few more attacks (AT=1 vs. AT=3 in Figure 12b), but their penalization is inevitable and they will be quickly suppressed (Figure 12c). Furthermore, the reputation engine can choose a higher initial penalty value in the system setup (i.e., p^(1) in Equation 1). This will also boost the process of penalizing faulty servers.

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