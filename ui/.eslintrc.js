module.exports = {
  extends: [
    "universe",
    "universe/shared/typescript-analysis",
    "plugin:react-hooks/recommended",
    "plugin:@typescript-eslint/recommended",
  ],
  plugins: [
    "unused-imports",
    "@typescript-eslint",
    "react-hooks",
    "react-refresh",
    "unused-imports",
  ],
  overrides: [
    {
      files: ["*.ts", "*.tsx", "*.d.ts"],

      parserOptions: {
        project: "./tsconfig.json",
      },
    },
  ],
  ignorePatterns: ["*.config.js", ".eslintrc.cjs", "*.config.ts", "*.d.ts", "**/node_modules/**", "*.eslintrc*"],
  settings: {
    "import/resolver": {
      typescript: {}, // this loads <rootdir>/tsconfig.json to ESLint
    },
  },
  /* for lint-staged */
  globals: {
    __dirname: true,
  },
  rules: {
    "react-refresh/only-export-components": [
      "warn",
      { allowConstantExport: true },
    ],
    "no-console": "error",
    "unused-imports/no-unused-imports": "error",
    "no-unused-vars": "off",
    "@typescript-eslint/no-explicit-any": "off",
  },
};
