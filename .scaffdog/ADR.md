---
name: 'ADR'
root: '.'
output: './docs/ADR'
ignore: ['.']
questions:
  number: 'バージョン管理'
  category: '何の技術選定をしたか教えてね'
  fileName: '決定した技術を教えてね'

---

# `{{ inputs.number | rtrim }}_{{ inputs.fileName }}.md`

```markdown
# ADR (Architecture Decision Record) for {{ inputs.fileName }}

## {{ inputs.category }}の技術選定

### 背景

### 決定

### 理由

### 代替案

### 影響

### 参考情報

執筆日：{{ date "YYYY/MM/DD HH:mm" }}

```