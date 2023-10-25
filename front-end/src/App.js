import React, { useState, useEffect } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import Header from './components/Header';

import Home from './components/Home';
import Error from './components/Error';
import ListRacer from './components/ListRacer';
import EditRacer from './components/EditRacer';

function App() {
    return (
        <BrowserRouter>
            <Header />
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/*" element={<Error />} />
                <Route path="/racers" element={<ListRacer />} />
                <Route path="/edit/:id" element={<EditRacer />} />
            </Routes>
        </BrowserRouter>
    );
}

export default App;
