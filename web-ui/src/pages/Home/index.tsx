import React, { useState, useEffect } from 'react';
import type { LogEntry, LogQuery } from '../../types';
import { apiService } from '../../services/api';

interface LogViewerProps {
  projectId: number;
}

const LogViewer: React.FC<LogViewerProps> = ({ projectId = 1 }) => {
  const [logs, setLogs] = useState<LogEntry[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [query, setQuery] = useState<LogQuery>({});

  const fetchLogs = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await apiService.getProjectLogs(projectId, query);
      setLogs(response.logs);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch logs');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchLogs();
  }, [projectId, query]);

  const getLogLevelColor = (level: LogEntry['level']) => {
    switch (level) {
      case 'error':
        return 'text-red-600 bg-red-50';
      case 'warn':
        return 'text-yellow-600 bg-yellow-50';
      case 'info':
        return 'text-blue-600 bg-blue-50';
      case 'debug':
        return 'text-gray-600 bg-gray-50';
      default:
        return 'text-gray-600 bg-gray-50';
    }
  };

  const formatTimestamp = (timestamp: string) => {
    return new Date(timestamp).toLocaleString();
  };

  return (
    <div className="log-viewer">
      {/* Search Controls */}
      <div className="mb-4 flex gap-4">
        <input
          type="text"
          placeholder="Search logs..."
          className="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          value={query.query || ''}
          onChange={(e) => setQuery({ ...query, query: e.target.value })}
          onKeyPress={(e) => e.key === 'Enter' && fetchLogs()}
        />
        <button
          onClick={fetchLogs}
          disabled={loading}
          className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:opacity-50"
        >
          {loading ? 'Searching...' : 'Search'}
        </button>
      </div>

      {/* Error Message */}
      {error && (
        <div className="mb-4 p-3 bg-red-50 border border-red-200 rounded-md">
          <p className="text-red-600">{error}</p>
        </div>
      )}

      {/* Logs Display */}
      <div className="bg-gray-50 border rounded-md max-h-96 overflow-y-auto">
        {logs.length === 0 && !loading ? (
          <div className="p-4 text-center text-gray-500">No logs found</div>
        ) : (
          <div className="divide-y divide-gray-200">
            {logs.map((log, index) => (
              <div key={index} className="p-3 hover:bg-gray-100">
                <div className="flex items-start gap-3">
                  <span className={`px-2 py-1 text-xs font-medium rounded ${getLogLevelColor(log.level)}`}>
                    {log.level.toUpperCase()}
                  </span>
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center gap-2 text-sm text-gray-500 mb-1">
                      <span>{formatTimestamp(log.timestamp)}</span>
                      <span>•</span>
                      <span>{log.source}</span>
                    </div>
                    <div className="text-gray-900 break-words">{log.message}</div>
                    {log.fields && Object.keys(log.fields).length > 0 && (
                      <div className="mt-2 text-xs text-gray-600">
                        <details>
                          <summary className="cursor-pointer hover:text-gray-800">
                            Fields ({Object.keys(log.fields).length})
                          </summary>
                          <pre className="mt-1 bg-gray-100 p-2 rounded overflow-x-auto">
                            {JSON.stringify(log.fields, null, 2)}
                          </pre>
                        </details>
                      </div>
                    )}
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Loading Indicator */}
      {loading && (
        <div className="mt-4 text-center">
          <div className="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600"></div>
          <span className="ml-2 text-gray-600">Loading logs...</span>
        </div>
      )}
    </div>
  );
};

export default LogViewer;
