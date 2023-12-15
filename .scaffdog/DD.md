---
name: 'DD'
root: '.'
output: './docs/DesignDoc'
ignore: ['.']
questions:
  number: 'バージョン管理'
  fileName: '何の機能のDDか教えてね'

---

# `{{ inputs.number | rtrim }}_{{ inputs.fileName }}.md`

```markdown
# DesignDoc

## {{ inputs.fileName }}機能について

### 概要

### 詳細

### シーケンス図

### ER図

### スキーマ

執筆日：{{ date "YYYY/MM/DD HH:mm" }}

```