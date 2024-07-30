import { RegistrationPage } from '@pages/index';
import { useDispatch } from 'react-redux';
import { Route, Routes } from 'react-router-dom';
import 'remixicon/fonts/remixicon.css';
import './App.scss';
import store from './store';

export default function App(): JSX.Element {
  return (
    <div className='h-full'>
      <Routes>
        <Route path='/' element={<RegistrationPage />} />
      </Routes>
    </div>
  );
}

export type AppDispatch = typeof store.dispatch;

export const useAppDispatch = () => useDispatch<AppDispatch>();