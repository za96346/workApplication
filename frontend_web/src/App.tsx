import React from 'react';

import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
  useNavigate,
} from 'react-router-dom';
// @ts-ignore,
import Layout from './page/Layout.tsx';
// @ts-ignore,
import TimeSettingPage from './page/TimeSettingPage';
import EditPage from './page/EditPage';
import NoticePage from './page/NoticePage';


function App() {
    return (
      	<Router>
            <Routes>
                <Route element={<Layout />}>
                    <Route path='/' element={<NoticePage />} />
                    <Route path='edit' element={<EditPage />} />
                    <Route path='timeSetting' element={<TimeSettingPage />} />
                </Route>
            </Routes>
        </Router>
    );
}

export default App;
