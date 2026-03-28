import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Layout } from "./components/Layout";
import { ProtectedRoute } from "./components/ProtectedRoute";
import { AuthProvider } from "./context/AuthContext";
import { AddTaskPage } from "./pages/AddTaskPage";
import { AuthPage } from "./pages/AuthPage";
import { DashboardPage } from "./pages/DashboardPage";
import { ProgressPage } from "./pages/ProgressPage";
import { TimetablePage } from "./pages/TimetablePage";

export default function App() {
  return (
    <AuthProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/auth" element={<AuthPage />} />
          <Route
            path="/"
            element={
              <ProtectedRoute>
                <Layout>
                  <DashboardPage />
                </Layout>
              </ProtectedRoute>
            }
          />
          <Route
            path="/add-task"
            element={
              <ProtectedRoute>
                <Layout>
                  <AddTaskPage />
                </Layout>
              </ProtectedRoute>
            }
          />
          <Route
            path="/progress"
            element={
              <ProtectedRoute>
                <Layout>
                  <ProgressPage />
                </Layout>
              </ProtectedRoute>
            }
          />
          <Route
            path="/timetable"
            element={
              <ProtectedRoute>
                <Layout>
                  <TimetablePage />
                </Layout>
              </ProtectedRoute>
            }
          />
        </Routes>
      </BrowserRouter>
    </AuthProvider>
  );
}

