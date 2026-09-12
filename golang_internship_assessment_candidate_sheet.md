# 📝 Backend Engineering Internship Assessment

**Candidate Name:** Habibina Arif Muzayyan  
**University / Major:** Institut Teknologi Bandung / Teknik Informatika
**Duration:**  (2 Hours)  

---

## 📢 Pre-Assessment Notice & Recommended Study Topics (Pemberitahuan Sebelum Mulai)

> [!IMPORTANT]
> **Dear Candidate,**  
> Before starting this technical assessment, please ensure you have reviewed and familiarized yourself with the following core topics:
>
> 1. **Git & Version Control Workflow:**
>    - Basic commands: `git init`, `git add`, `git commit`, `git push`, `git status`, `git branch`, `git stash`.
>    - Differences between `git merge` and `git rebase`.
>    - Resolving Git merge conflicts and amending unpushed commits (`git commit --amend`).
>    - Writing clear, descriptive commit messages.
>
> 2. **Golang Fundamentals:**
>    - Core data types, zero-values (`nil`, `0`, `""`), slices & underlying arrays (`len` vs `cap`), and maps (safe `comma-ok` key lookups).
>    - Memory mechanics: pointers (`*Type`) vs pass-by-value, struct definitions, and pointer receivers.
>    - Control flow & idioms: `defer` LIFO execution, explicit error handling (`if err != nil`), and implicit interfaces.
>    - JSON serialization and struct tags (e.g., `json:"field_name"`, `omitempty`, `json:"-"`).
>
> 3. **Databases (MySQL & MongoDB):**
>    - **MySQL (Relational):** Basic SQL queries, `INNER JOIN` vs `LEFT JOIN`, `PRIMARY KEY` vs `UNIQUE` key, indexes, and ACID transactions.
>    - **MongoDB (NoSQL Document):** Collections vs documents, BSON structure, nested sub-documents, basic `find()` queries, and `$lookup` aggregations.
>    - **Database Security:** Preventing SQL Injection using parameterized queries and proper resource cleanup.

---

## 🚀 Submission & Git Workflow Instructions (MANDATORY)

To evaluate your practical Git workflow and version control habits, you are required to submit your work via **GitHub**:

1. **Create a Public GitHub Repository:**
   - Create a new **public** repository on your personal GitHub account named `backend-intern-assessment` (or `intern-assessment-<your-name>`).
   - Clone or initialize the repository on your local computer.
   - Copy this assessment markdown file (or create your answer files) inside the repository root.

2. **Section-by-Section Commit & Push Rule:**
   - **Do NOT submit all answers in a single commit!**
   - Each time you complete a section, create a dedicated Git commit with a descriptive commit message and push it immediately to GitHub before moving to the next section:
     - 📌 **After Section 1:** `git add . --> git commit -m "feat: complete section 1 - single choice mcq" --> git push`
     - 📌 **After Section 2:** `git add . --> git commit -m "feat: complete section 2 - multi-select checkbox" --> git push`
     - 📌 **After Section 3:** `git add . --> git commit -m "feat: complete section 3 - short essay questions" --> git push`
     - 📌 **After Section 4:** Put your Go solution in `main.go` (or inside this markdown) --> `git commit -m "feat: complete section 4 - voucher engine coding problem"` --> `git push`
     - 📌 **Final Submission:** Review your answers --> `git commit -m "chore: finalize assessment submission"` --> `git push origin main`

3. **Language Requirement:** All answers, code comments, and commit messages must be in **English**.

---

## ✍️ Quick Answer Sheet (Sections 1 & 2)

### Section 1: Single-Choice (Write A, B, C, or D)
| Q1 | Q2 | Q3 | Q4 | Q5 | Q6 | Q7 | Q8 | Q9 | Q10 | Q11 | Q12 | Q13 | Q14 | Q15 |
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] | [ &nbsp; ] |

### Section 2: Multiple Select / Checkbox (Write all correct letters, e.g., "A, B, D")
| Q16 | Q17 | Q18 | Q19 | Q20 |
|:---:|:---:|:---:|:---:|:---:|
| [ &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; ] | [ &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; ] | [ &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; ] | [ &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; ] | [ &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; ] |

---

# SECTION 1: Single-Choice Multiple Choice
*Instructions: Choose the single best answer for each question (A, B, C, or D).*  
*👉 Commit & push after completing this section.*

### Part A: Git & Version Control

**1. You modified multiple files, but only want to stage and commit one specific file (`auth.go`). Which command sequence is correct?**
- A) `git commit -a -m "add auth"`
- B) `git add auth.go` followed by `git commit -m "feat: implement auth logic"`
- C) `git push origin auth.go` followed by `git commit`
- D) `git checkout -b auth.go` followed by `git commit`

**2. What is the key operational difference between `git merge` and `git rebase` when integrating changes from `main` into your feature branch?**
- A) `git merge` deletes the feature branch history, while `git rebase` preserves it.
- B) `git merge` creates a new merge commit preserving exact chronological branch history; `git rebase` rewrites commit history onto the tip of `main` for a linear history.
- C) `git rebase` can only be performed on the remote repository.
- D) `git merge` is only used for pulling remote changes; `git rebase` is used only for stashing.

**3. You committed code with the message `"fix typo"`, but you forgot to include a modified configuration file (`config.yaml`). If the commit has NOT been pushed to remote yet, what is the cleanest way to include the file into the previous commit?**
- A) Run `git reset --hard HEAD~1` and start over.
- B) Stage the config file (`git add config.yaml`) and run `git commit --amend --no-edit`.
- C) Delete the `.git` folder and reinitialize.
- D) Create a new branch and push both commits separately.

**4. What happens when a Git merge conflict occurs?**
- A) Git automatically deletes conflicting files and retains remote files.
- B) Git rolls back the repository to the initial commit.
- C) Git marks the conflicting sections in the affected files with `<<<<<<<`, `=======`, and `>>>>>>>` markers, pausing the merge until resolved and committed.
- D) Git prevents any further commits until the remote repository is restarted.

---

### Part B: Golang Fundamentals

**5. Consider the following Go code snippet:**
```go
package main
import "fmt"

func modifySlice(s []int) {
    s[0] = 99
    s = append(s, 100)
}

func main() {
    nums := []int{1, 2, 3}
    modifySlice(nums)
    fmt.Println(nums)
}
```
**What will be printed to the console?**
- A) `[99 2 3 100]`
- B) `[1 2 3]`
- C) `[99 2 3]`
- D) `[99 100 3]`

**6. In Go, what is the default zero value of a pointer, a slice, a map, and an interface?**
- A) `0`
- B) `nil`
- C) `undefined`
- D) Empty struct `{}`

**7. How are errors typically handled in idiomatic Go?**
- A) By wrapping code in `try...catch...finally` blocks.
- B) By returning `(result, error)` as multiple return values and checking `if err != nil`.
- C) By using `panic()` and `recover()` for standard validation errors.
- D) By checking global runtime error variables.

**8. What is the behavior of the `defer` statement in the following code?**
```go
package main
import "fmt"

func count() {
    for i := 1; i <= 3; i++ {
        defer fmt.Print(i, " ")
    }
}

func main() {
    count()
}
```
- A) `1 2 3 `
- B) `3 2 1 `
- C) `3 3 3 `
- D) `1 1 1 `

**9. Why is passing a large struct by pointer (`*User`) instead of by value (`User`) often preferred for functions that modify the user data?**
- A) Because Go does not support methods on struct values.
- B) Because passing by pointer avoids copying the entire struct in memory and allows the function to mutate the original struct's fields directly.
- C) Because value types cannot be stored in databases.
- D) Because Go's compiler will throw a syntax error if pointers are not used.

**10. In Go, how do you safely check whether a specific key exists in a map `userMap := map[string]int{"andi": 10}`?**
- A) `if userMap.contains("andi") { ... }`
- B) `if val, exists := userMap["andi"]; exists { ... }`
- C) `if userMap.hasKey("andi") { ... }`
- D) `if userMap["andi"] != undefined { ... }`

---

### Part C: MySQL & MongoDB Databases

**11. In MySQL (InnoDB), what is the difference between a `PRIMARY KEY` and a `UNIQUE` key?**
- A) A table can have multiple `PRIMARY KEY`s, but only one `UNIQUE` key.
- B) A `PRIMARY KEY` cannot contain `NULL` values and creates a clustered index; a `UNIQUE` key enforces uniqueness but may permit `NULL` values.
- C) `UNIQUE` keys are only supported in NoSQL databases.
- D) `PRIMARY KEY` only works on integer data types.

**12. Which SQL statement retrieves all users who registered in 2026, ordering the newest registrations first, limiting the results to 10 rows?**
- A) `SELECT * FROM users WHERE YEAR(created_at) = 2026 ORDER BY created_at DESC LIMIT 10;`
- B) `SELECT * FROM users HAVING created_at = 2026 SORT BY created_at TOP 10;`
- C) `GET ALL FROM users FILTER BY 2026 ORDER DESCENDING LIMIT 10;`
- D) `SELECT TOP 10 * FROM users WHERE created_at IN (2026) GROUP BY created_at;`

**13. What is the fundamental difference between an `INNER JOIN` and a `LEFT JOIN` in SQL?**
- A) `INNER JOIN` returns only rows with matching keys in both tables; `LEFT JOIN` returns all rows from the left table and matched rows from the right table (with `NULL` for non-matches).
- B) `INNER JOIN` is for MongoDB, while `LEFT JOIN` is for MySQL.
- C) `LEFT JOIN` deletes rows without foreign key references.
- D) `INNER JOIN` is always executed in memory, while `LEFT JOIN` is executed on disk.

**14. What is the purpose of database transactions and the "ACID" property?**
- A) To compress SQL query responses to reduce network latency.
- B) To ensure operations are Atomicity, Consistency, Isolation, and Durability guaranteed, preventing partial updates during failures.
- C) To allow multiple users to bypass password authentication.
- D) To automatically convert SQL tables into MongoDB collections.

**15. In MongoDB, what is the equivalent of a SQL "Table" and a SQL "Row"?**
- A) Database and Column
- B) Collection and Document
- C) Document and Field
- D) Index and Aggregation

---

# SECTION 2: Multiple-Select / Checkbox Questions
*Instructions: Select **ALL** correct options for each question. There may be 2, 3, or more correct choices per question.*  
*👉 Commit & push after completing this section.*

**16. Which of the following statements about Git commands and everyday workflows are TRUE? (Select ALL that apply)**
- [ ] A) `git branch -d feature-branch` deletes a local branch.
- [ ] B) `git stash` temporarily shelves uncommitted local changes so you can work on a clean directory.
- [ ] C) `git push origin main` automatically deletes all remote branches that were merged locally.
- [ ] D) `git status` displays the state of the working directory and the staging area.
- [ ] E) `git init` converts an existing directory into a new Git repository.

**17. Which of the following types in Go have `nil` as their default zero value? (Select ALL that apply)**
- [ ] A) Slices (`[]int`)
- [ ] B) Maps (`map[string]string`)
- [ ] C) Pointers (`*User`)
- [ ] D) Integers (`int`)
- [ ] E) Interfaces (`error` or `any`)

**18. Which of the following statements about Go data structures, memory, and syntax are CORRECT? (Select ALL that apply)**
- [ ] A) Attempting to write into an uninitialized `nil` map (`var m map[string]int; m["key"] = 1`) triggers a fatal runtime panic.
- [ ] B) A Go slice header contains three fields: a pointer to the underlying array, length (`len`), and capacity (`cap`).
- [ ] C) Struct tags such as `` `json:"user_id,omitempty"` `` allow customizing JSON field names and omitting empty fields during serialization.
- [ ] D) In Go, a struct must explicitly declare an `implements` keyword to satisfy an interface.
- [ ] E) The `make()` built-in function is used to initialize slices, maps, and channels with allocated memory.

**19. Which of the following practices are recommended for building secure, reliable, and performant database applications? (Select ALL that apply)**
- [ ] A) Use parameterized SQL queries (e.g., `db.Query("SELECT ... WHERE email = ?", email)`) to prevent SQL Injection attacks.
- [ ] B) Embed raw user input strings directly into SQL queries using `fmt.Sprintf` to maximize query execution speed.
- [ ] C) Add indexes to columns that are frequently used in `WHERE` filters, `JOIN` conditions, and `ORDER BY` clauses.
- [ ] D) Always close SQL query result sets (`rows.Close()`) to release database connections back to the connection pool.
- [ ] E) Use database transactions (`BEGIN`, `COMMIT`, `ROLLBACK`) when executing multi-step financial balance transfers.

**20. Which of the following statements regarding MongoDB (NoSQL) are TRUE? (Select ALL that apply)**
- [ ] A) MongoDB stores records as flexible, semi-structured BSON (Binary JSON) documents.
- [ ] B) Different documents within the same MongoDB collection can contain completely different fields and data structures.
- [ ] C) The `$lookup` aggregation stage allows joining data from another collection (similar to a SQL LEFT OUTER JOIN).
- [ ] D) MongoDB does not support creating indexes on document fields.
- [ ] E) MongoDB natively supports nested sub-documents and arrays within a single document.

---

# SECTION 3: Short Essay Questions
*👉 Commit & push after completing this section.*

### Question 3.1: Go Struct Tags & JSON Security
**Scenario:**  
You are building a REST API in Go. You have a `User` struct that you want to return to frontend clients via JSON:
```go
type User struct {
    ID           int
    FullName     string
    Email        string
    PasswordHash string
}
```
However, two requirements must be fulfilled:
1. Frontend requires JSON keys in lowercase snake_case (e.g., `"full_name"` and `"email"`).
2. The `PasswordHash` field is sensitive security data and must **never** be included in the JSON output sent to the client.

- **Task:**
  1. Rewrite the `User` struct using Go struct tags (`json:"..."`) to satisfy both requirements.
  2. Explain what the `json:"-"` struct tag does during JSON serialization (`json.Marshal`).

*Write your answer here:*
```text


```

---

### Question 3.2: Database Selection (MySQL vs. MongoDB)
**Scenario:**  
Your team is building an **E-Commerce Platform**. You need to store two main modules:
1. **User Wallets & Financial Balances:** Requires atomic transfers, zero data loss, strict consistency, and transaction rollbacks on failure.
2. **Product Catalog & Dynamic Specifications:** Products vary widely in attributes (e.g., Laptops have CPU/RAM, T-Shirts have Size/Color/Fabric, Shoes have Shoe Sizes).

- **Task:**
  - Which database (MySQL or MongoDB) would you choose for **User Wallets**, and which for **Product Catalog**?
  - Briefly justify your choices based on schema flexibility and ACID transaction guarantees.

*Write your answer here:*
```text


```

---

# SECTION 4: Story Problem - Program Design & Coding
*👉 Implement your code, then make your final commit & push.*

### Scenario: "FlashSale Voucher Discount Engine"
You are tasked with building a core discount calculation component for an Indonesian e-commerce platform in **Golang**.

#### Requirements:
1. Define a Go `struct` named `CartItem` with fields:
   - `ProductID` (`string`)
   - `Name` (`string`)
   - `Price` (`float64`)
   - `Quantity` (`int`)
2. Define a Go `struct` named `Voucher` with fields:
   - `Code` (`string`)
   - `DiscountPercent` (`float64`, e.g., `10.0` for 10%)
   - `MaxDiscount` (`float64`, e.g., `50000.0` maximum discount cap)
   - `MinPurchase` (`float64`, minimum total cart amount required)
3. Implement a function with the following signature:
   ```go
   func CalculateFinalPrice(items []CartItem, voucher *Voucher) (subtotal float64, discount float64, total float64, err error)
   ```
4. **Business Validation Rules:**
   - If `items` slice is empty, return an error: `"cart cannot be empty"`.
   - If any item has `Quantity <= 0` or `Price < 0`, return an error: `"invalid item price or quantity"`.
   - Calculate `subtotal` as the sum of `Price * float64(Quantity)` for all items.
   - If `voucher` is `nil`, discount is `0`, and `total = subtotal`.
   - If `voucher` is provided:
     - Check if `subtotal >= voucher.MinPurchase`. If not, no discount is applied (`discount = 0`), but return **no error** (customer pays normal subtotal).
     - If eligible, calculate nominal discount = `subtotal * (voucher.DiscountPercent / 100)`.
     - If nominal discount exceeds `voucher.MaxDiscount`, cap the discount at `voucher.MaxDiscount`.
     - Final `total = subtotal - discount`.

- **Task:**
  Write the complete, runnable Go code implementing the structs, the `CalculateFinalPrice` function, and a simple `main()` test case.

*Write your Go code here:*
```go
package main

// TODO: Write your code here
```
