import React from 'react';
import { NavLink } from 'react-router-dom';

function Home() {
    return (
        <div className="App bg-green-700">
            <h1 className="title text-3xl font-bold text-center">Bienvenue sur le site de gestion de course de Cali</h1>
            <div className="buttons">
                <NavLink to="/racers" className="button mx-2 bg-green-600 py-1 px-2 rounded-lg">Liste des coureurs</NavLink>
                <NavLink to="/edit" className="button mx-2 bg-green-600 py-1 px-2 rounded-lg">Ajouter un coureur</NavLink>
            </div>
        </div>
    );
}

export default Home;
