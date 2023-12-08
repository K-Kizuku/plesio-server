module.exports = {
    disableEmoji: false,
    format: '{type}{scope}: {emoji}{subject}',
    list: ['feat', 'fix' ,'test', 'docs', 'setting', 'refactor', 'style', 'ci', 'perf'],
    maxMessageLength: 64,
    minMessageLength: 1,
    questions: ['type', 'scope', 'subject'],
    scopes: [],
    types: {
      ci: {
        description: 'CIに関する追加・修正',
        emoji: '🎡',
        value: 'ci'
      },
      docs: {
        description: 'ドキュメントの追加・修正',
        emoji: '📄',
        value: 'docs'
      },
      feat: {
        description: '機能実装',
        emoji: '🍻',
        value: 'feat'
      },
      fix: {
        description: 'バグ修正',
        emoji: '🐛',
        value: 'fix'
      },
      perf: {
        description: 'パフォーマンス改善',
        emoji: '🚀',
        value: 'perf'
      },
      refactor: {
        description: 'リファクタ',
        emoji: '💡',
        value: 'refactor'
      },
      style: {
        description: 'コーディングスタイルの修正',
        emoji: '💄',
        value: 'style'
      },
      test: {
        description: 'テストの追加・修正',
        emoji: '🧪',
        value: 'test'
      },
      setting:{
        description: '設定ファイルの追加・修正',
        emoji:'🔧',
        value: 'setting'
      },
    },
    messages: {
      type: 'どんなコミットか選んでね:',
      subject: '簡潔に変更点を教えてね\n',
      issues: 'issue番号を教えてね\n'
    }
  };