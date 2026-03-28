import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// Vite is the frontend development tool. It gives fast local development and builds.
export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173,
  },
});

