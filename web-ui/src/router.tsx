import { createBrowserRouter, redirect } from 'react-router';

import { ConsoleLayout } from './layouts/ConsoleLayout';
import Home from './pages/Home';
import Login from './pages/Login';

export const router = createBrowserRouter([
  { path: '/login', element: <Login /> },
  {
    path: '/',
    element: <ConsoleLayout />,
    children: [{ path: '/', element: <Home />, index: true }],
    loader: () => {
      // if (!sessionStorage.getItem('logged')) {
      //   throw redirect('/login');
      // }
    },
  },
]);
