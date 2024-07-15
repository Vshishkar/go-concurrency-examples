# go-concurrency-examples

This repo showcases simple way of implementing in-memory transactions in golang.
The main focus is going to be around different isolation levels and different concurrency controls.
The domain used for examples is in-memory banking system.

| Isolation level / Concurrency control  | Read uncommitted  | Read committed  | Snapshot  | Repeatable read | Serializable  |
|---|---|---|---|---|---|
| Optimistic    |   |   |   |   |   |
| Pessimistic   |   |   |   |   |   |
| Multi-version  |   |   |   |   |   |

Main idea is to simulate database as in-memory object and implement different strategies concurrency control strategies.
 