# Word Count Tool

## Description
This tool identifies the number of valid words in a given string. A word is considered valid if it meets all of the following criteria:

- Contains at least 3 characters
- Contains only alphanumeric characters (letters 'a'-'z', 'A'-'Z', and numbers 0-9)
- Includes at least one vowel ('a', 'e', 'i', 'o', 'u')
- Includes at least one consonant (any letter that is not a vowel)

The tool returns the count of valid words in the input string.

## Example
Suppose the input string is:

| Word    | Is Valid | Reason                                            |
|---------|-----------|--------------------------------------------------|
| This    | Yes       | At least 3 characters, contains a vowel and a consonant |
| is      | No        | Less than 3 characters                            |
| an      | No        | Less than 3 characters                            |
| example | Yes       | At least 3 characters, contains a vowel and a consonant |
| string  | Yes       | At least 3 characters, contains a vowel and a consonant |
| 234     | No        | Does not contain a vowel or a consonant          |

## Function Description
Complete the function `countValidWords` with the following parameter(s):

- `string s`: a string to analyze

### Returns
- `int`: the number of valid words in `s`

### Constraints
- 1 ≤ |s| ≤ 10⁵
- `s` consists of all available ASCII characters.

## Input Format For Custom Testing
Input will be provided as a string `s`.

## Sample Case

**Sample Input**
- This is Form16 submis$ion date
  
**Sample Output**

**Explanation**

Only 'This', 'Form16', and 'date' are valid words. Since 'is' only contains 2 characters and 'submis$ion' has an invalid character, they are not valid.

---

