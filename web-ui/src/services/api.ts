import axios, { type AxiosInstance } from 'axios';
import type {
  User,
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  Project,
  CreateProjectRequest,
  LogResponse,
  LogQuery,
  RetentionPolicy,
  UpdateRetentionRequest,
} from '../types';

class ApiService {
  private client: AxiosInstance;

  constructor() {
    this.client = axios.create({
      baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:9012/api/v1',
      timeout: 10000,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Add request interceptor to include auth token
    this.client.interceptors.request.use((config) => {
      const token = localStorage.getItem('auth_token');
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    });

    // Add response interceptor for error handling
    this.client.interceptors.response.use(
      (response) => response,
      (error) => {
        if (error.response?.status === 401) {
          localStorage.removeItem('auth_token');
          window.location.href = '/login';
        }
        return Promise.reject(error);
      }
    );
  }

  // Auth methods
  async login(credentials: LoginRequest): Promise<LoginResponse> {
    const response = await this.client.post<LoginResponse>('/auth/login', credentials);
    return response.data;
  }

  async register(userData: RegisterRequest): Promise<{ message: string; user: User }> {
    const response = await this.client.post('/auth/register', userData);
    return response.data;
  }

  // User methods
  async getUsers(): Promise<User[]> {
    const response = await this.client.get<{ users: User[] }>('/users');
    return response.data.users;
  }

  async createUser(userData: Omit<User, 'id' | 'created_at' | 'updated_at'>): Promise<User> {
    const response = await this.client.post('/users', userData);
    return response.data.user;
  }

  // Project methods
  async getProjects(): Promise<Project[]> {
    const response = await this.client.get<{ projects: Project[] }>('/projects');
    return response.data.projects;
  }

  async createProject(projectData: CreateProjectRequest): Promise<Project> {
    const response = await this.client.post('/projects', projectData);
    return response.data.project;
  }

  async getProject(id: number): Promise<Project> {
    const response = await this.client.get<{ project: Project }>(`/projects/${id}`);
    return response.data.project;
  }

  async updateProject(id: number, projectData: Partial<CreateProjectRequest>): Promise<Project> {
    const response = await this.client.put(`/projects/${id}`, projectData);
    return response.data.project;
  }

  async deleteProject(id: number): Promise<void> {
    await this.client.delete(`/projects/${id}`);
  }

  // Log methods
  async getProjectLogs(projectId: number, query?: LogQuery): Promise<LogResponse> {
    const params = new URLSearchParams();
    if (query?.query) params.append('query', query.query);
    if (query?.start) params.append('start', query.start);
    if (query?.end) params.append('end', query.end);

    const response = await this.client.get<LogResponse>(
      `/projects/${projectId}/logs?${params.toString()}`
    );
    return response.data;
  }

  // Configuration methods
  async getRetentionPolicies(): Promise<RetentionPolicy[]> {
    const response = await this.client.get<{ policies: RetentionPolicy[] }>('/config/retention');
    return response.data.policies;
  }

  async updateRetentionPolicy(policyData: UpdateRetentionRequest): Promise<RetentionPolicy> {
    const response = await this.client.put('/config/retention', policyData);
    return response.data.policy;
  }

  // Health check
  async healthCheck(): Promise<{ status: string; service: string }> {
    const response = await this.client.get('/health');
    return response.data;
  }
}

// Export singleton instance
export const apiService = new ApiService();
export default apiService;
