--- :LocalSettingsExample & :LocalSettingsReload
return {
  lsp = {
    --- https://github.com/golang/tools/blob/master/gopls/doc/settings.md
    gopls = {
      semanticTokens = true,  -- 默认 false.
      usePlaceholders = false,
    },
    -- gopls = { ["ui.completion.usePlaceholders"] = true },  --- 两种写法都成立
  },
  -- linter = {
  --   golangci_lint = {
  --     extra_args = { "-c", ".golangci.yml" }
  --   },
  -- }
}
