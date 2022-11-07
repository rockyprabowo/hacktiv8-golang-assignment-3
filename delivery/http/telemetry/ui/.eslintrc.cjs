module.exports = {
    env: {
        browser: true,
        es2021: true
    },
    extends: [
        'plugin:react/jsx-runtime',
        'plugin:react/recommended',
        'standard-with-typescript',
        'plugin:@typescript-eslint/recommended'
    ],
    overrides: [],
    parserOptions: {
        ecmaVersion: 'latest',
        sourceType: 'module',
        ecmaFeatures: {
            jsx: true
        },
        project: ['./tsconfig.json']
    },
    plugins: [
        'react-hooks',
        'react',
        '@typescript-eslint'
    ],
    rules: {
        quotes: ['error', 'single'],
        'linebreak-style': ['error', 'unix'],
        'no-empty': 'warn',
        'no-console': 'warn',
        '@typescript-eslint/no-use-before-define': ['error'],
        '@typescript-eslint/no-unused-vars': [
            'error',
            {vars: 'all', args: 'after-used', ignoreRestSiblings: false}
        ],
        '@typescript-eslint/no-shadow': ['error'],
        '@typescript-eslint/explicit-function-return-type': 'warn', // Consider using explicit annotations for object literals and function return types even when they can be inferred.
        'react/jsx-filename-extension': ['warn', {extensions: ['.tsx']}],
        'react/react-in-jsx-scope': 'off',
        'import/extensions': [
            'error',
            'ignorePackages',
            {
                ts: 'never',
                tsx: 'never'
            }
        ],
        'react-hooks/rules-of-hooks': 'error',
        'react-hooks/exhaustive-deps': 'warn'
    },
    settings: {
        'import/resolver': {
            typescript: {}
        },
        react: {
            version: 'detect'
        }
    }
}
