# To-Do: Future Improvements

- [ ] **Move to the native React Compiler** — replace the Babel-based compiler (`@rolldown/plugin-babel` + `babel-plugin-react-compiler` + `reactCompilerPreset` in `frontend/vite.config.ts`) with `@vitejs/plugin-react`'s Rust compiler (`react({ compiler: true })` + `oxc-transform-react`) once it's no longer experimental. Removes Babel from the build entirely.
- [ ] **Add GitHub actions**
