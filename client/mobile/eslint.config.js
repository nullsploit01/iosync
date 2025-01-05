import { parse as tsParser } from '@typescript-eslint/parser'
import reactPlugin from 'eslint-plugin-react'
import unusedImports from 'eslint-plugin-unused-imports'

export default [
  {
    // Specify the parser for TypeScript
    parser: tsParser,
    parserOptions: {
      ecmaVersion: 2021,
      sourceType: 'module',
      ecmaFeatures: {
        jsx: true // Enable JSX parsing
      },
      project: './tsconfig.json' // Specify your TypeScript config
    },
    plugins: {
      'unused-imports': unusedImports,
      react: reactPlugin
    },
    overrides: [
      {
        files: ['*.ts', '*.tsx', '*.d.ts'],

        parserOptions: {
          project: './tsconfig.json'
        }
      }
    ],
    ignorePatterns: ['*.config.js', '.eslintrc.cjs', '*.config.ts'],
    settings: {
      'import/resolver': {
        typescript: {} // this loads <rootdir>/tsconfig.json to ESLint
      }
    },
    /* for lint-staged */
    globals: {
      __dirname: true
    },
    rules: {
      // Rules for unused imports and variables
      'no-console': 'error',
      'unused-imports/no-unused-imports': 'warn',
      'unused-imports/no-unused-vars': [
        'warn',
        {
          vars: 'all',
          varsIgnorePattern: '^_',
          args: 'after-used',
          argsIgnorePattern: '^_'
        }
      ],
      '@typescript-eslint/no-unused-vars': 'off', // Turn off in favor of unused-imports

      // React-specific rules
      'react/jsx-uses-react': 'warn',
      'react/jsx-uses-vars': 'warn'
    },
    settings: {
      react: {
        version: 'detect' // Automatically detect React version
      }
    }
  }
]
