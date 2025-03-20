## Transaction in SQL

  - What is Transaction in SQL?
  - 
    `Sequence of one or more SQL operations that are executed as a single unit of work.`

  - What is ACID properties?
  - 
    `ATOMATIC : Transaction is fully completed or fully Roll back`
    `Consistency : Database remain in a consistent state before and after the transaction.`
    `Isolation :  Transactions is do not interfere with each other; intermediate states are not visible.`
    `Durability : Once a transaction is committed, changes are permanent, even in case of a system failure.`

## Transaction Control Commands in SQL :

    BEGIN TRANSACTION: Starts a new transaction.
    COMMIT: Saves all changes made in the transaction.
    ROLLBACK: Undoes changes if an error occurs.
    SAVEPOINT: Creates a savepoint within a transaction for partial rollbacks.

## Running The Example :

  `Go run main.go`

- After Running the go file in MYSQL terminal using a Transaction Control Commands :

  ```
  BEGIN TRANSACTION;
  UPDATE history SET Balance = 5500.0 WHERE AccID = 2;
  COMMIT;
  ```

  `Change The Balance of AccID 2 and save all changes using a commit commands.`

  
