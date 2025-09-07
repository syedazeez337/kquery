param(
  [string]$Grammar = "grammar\KQuery.g4",
  [string]$Out = "runtime-go\parser"
)

New-Item -ItemType Directory -Force -Path $Out | Out-Null

# Generate Go lexer+parser
& antlr4 '-Dlanguage=Go' '-o' $Out '-package' 'parser' $Grammar
