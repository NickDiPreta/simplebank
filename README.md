# simplebank

This project is a simple banking app that uses a microservice architecture and Go on the backend.

&nbsp;

*The following principles were followed during DB design.*

- Atomicity: Either all operations complete successfully or the transaction fails and db is unchanged.
- Consistency: The DB state must be valid after the transaction and all constraints satisfied.
- Isolation: Concurrent transactions must not affect each other.
- Durability: Data written by a successful transaction must be recorded in persistent storage.

This was done to avoid the following read phenomena:

- Dirty reads caused by read-uncommitted isolation level
- Non-repeatable reads
- Phantom reads
- Serialization Anomaly
