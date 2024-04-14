# DataReconciler

## Intro: This project explores how transactional data replicated to a cloud data store (mongodb) could be reconciled against the source datastore.

The project exploits the use of mongomock https://pkg.go.dev/github.com/mjarkk/mongomock
and diff https://pkg.go.dev/github.com/r3labs/diff


This library will 

- populateSourceDB : populate a mocked source database (csv file!) with synthetic data
- populateTargetDB : replicate this data to the mongomock target database
- corruptTargetDB : mess up a small sample of this data to give us something that fails during reconciliation
- reconcileRecords : reconcile a set of records using data.diff which outputs a json representation of data differences. 
- These reconciliation json snippets are then put into a different collection on mongodb

```
{
    "TransactionID": "30c8adf3-1991-490a-96f2-da43cdfe619f",
    "Changelog": [
        {
            "Type": "update",
            "Field": "Reference",
            "From": "REF 8809",
            "To": "REF 8809 CORRUPT"
        },
        {
            "Type": "update",
            "Field": "Amount",
            "From": 664.67,
            "To": 0
        }
    ]
}

```