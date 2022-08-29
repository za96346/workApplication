import React from 'react';

import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
  useNavigate,
} from 'react-router-dom';
import styles from "./index.module.scss";

import Layout from './page/Layout';
import ListSettingPage from './page/ListSettingPage';
import EditPage from './page/EditPage';
import NoticePage from './page/NoticePage';
import SelfSettingPage from './page/SelfSettingPage';
import ListSearchPage from './page/ListSearchPage';
import language from './language';
import ErrorPage from './page/ErrorPage';


// global init
window.styles = styles;
window.language = language;

function App() {
    console.log(window.language)
    return (
      	<Router>
            <Routes>
                <Route element={<Layout />}>
                    <Route path='/' element={<NoticePage />} />
                    <Route path='edit' element={<EditPage />} />
                    <Route path='listSetting' element={<ListSettingPage />} />
                    <Route path='selfSetting' element={<SelfSettingPage />} />
                    <Route path='listSearch' element={<ListSearchPage />} />
                </Route>
                <Route path='*' element={<ErrorPage/>} />
            </Routes>
        </Router>
    );
}

export default App;
