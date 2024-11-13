import { Route, Routes } from 'react-router-dom';
import './App.css';
import Home from './pages/Home';
import { Plan } from './pages/Plan';


const App = () => {
  return (
   <Routes>
      <Route path='/' element={<Home />} />
      <Route path='/plan' element={<Plan />} />
   </Routes>
  );
};

export default App;
