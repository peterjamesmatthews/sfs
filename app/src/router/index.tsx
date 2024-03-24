import { createBrowserRouter } from 'react-router-dom';
import Node from '../components/Node';

export default createBrowserRouter([
  {
    path: "*",
    element: <Node  />,
  },
]);
