# 🧠 My Daily Go Assignment Workflow

Welcome to my Go learning journey! This repository contains my daily assignments from platforms like [zop.dev](https://zop.dev), where I practice Go concepts through structured exercises and real-world problem solving.

---

## 📁 Workflow Strategy

To keep my codebase clean, organized, and easy to manage, I follow a **branch-based Git workflow**:

### 🔧 Step-by-Step Workflow

1. **Start a new assignment**
    - Create a new branch for the assignment:
      ```bash
      git checkout -b assignment_<number>_<topic>
      ```
      ✅ Example:
      ```bash
      git checkout -b assignment_2_currency_converter
      ```

2. **Develop the solution**
    - Write clean, tested code for the assignment.
    - Keep functions modular and easy to understand.
    - Add comments and use meaningful variable names.

3. **Test the code thoroughly**
    - Run test cases if applicable.
    - Check for edge cases, formatting, and logical errors.

4. **Commit the work**
    - Use meaningful commit messages:
      ```bash
      git add .
      git commit -m "Complete assignment 2: Currency Converter with time-based greeting"
      ```

5. **Merge into master**
    - Switch back to the `master` branch and merge:
      ```bash
      git checkout master
      git merge assignment_2_currency_converter
      ```

6. **Push to GitHub (optional)**
   ```bash
   git push origin master
