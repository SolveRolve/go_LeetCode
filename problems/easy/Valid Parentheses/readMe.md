
# Valid Parentheses

## Problem Statement

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['`, and `']'`, determine if the input string is valid.

### The input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

### Examples:

**Example 1:**
```
Input: s = "()"
Output: true
```

**Example 2:**
```
Input: s = "()[]{}"
Output: true
```

**Example 3:**
```
Input: s = "(]"
Output: false
```

**Example 4:**
```
Input: s = "([)]"
Output: false
```

### Constraints:
- `1 <= s.length <= 10^4`
- `s` consists of parentheses only `'()[]{}'`.

---

## Approach

To solve this problem, we can use a **stack** to keep track of open brackets. Here's the approach:

1. Traverse through each character in the string:
    - If the character is an opening bracket (`'('`, `'{'`, `'['`), push it onto the stack.
    - If the character is a closing bracket (`')'`, `'}'`, `']'`), check the stack:
        - If the stack is empty or the top of the stack is not the matching open bracket, the string is invalid, return `false`.
        - Otherwise, pop the stack.

2. After processing the string, if the stack is empty, it means all the brackets were properly closed, so the string is valid, return `true`. If the stack still contains elements, return `false`.

---