// API Response types
export interface ApiResponse<T> {
  data?: T;
  error?: string;
  message?: string;
}

// User types
export interface User {
  id: number;
  username: string;
  email: string;
  role: 'admin' | 'user';
  created_at: string;
  updated_at: string;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}

export interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}

// Project types
export interface Project {
  id: number;
  name: string;
  description?: string;
  owner_id: number;
  created_at: string;
  updated_at: string;
}

export interface CreateProjectRequest {
  name: string;
  description?: string;
}

// Log types
export interface LogEntry {
  timestamp: string;
  level: 'debug' | 'info' | 'warn' | 'error';
  message: string;
  source: string;
  fields?: Record<string, unknown>;
}

export interface LogQuery {
  query?: string;
  start?: string;
  end?: string;
  limit?: number;
  offset?: number;
}

export interface LogResponse {
  logs: LogEntry[];
  total?: number;
  query?: string;
  start?: string;
  end?: string;
}

// Configuration types
export interface RetentionPolicy {
  id: number;
  project_id: number;
  duration_days: number;
  created_at: string;
  updated_at: string;
}

export interface UpdateRetentionRequest {
  project_id: number;
  duration_days: number;
}

// UI State types
export interface LoadingState {
  isLoading: boolean;
  error?: string;
}

export interface PaginationState {
  page: number;
  pageSize: number;
  total: number;
}
