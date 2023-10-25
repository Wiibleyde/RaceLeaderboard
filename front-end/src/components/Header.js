import React from 'react';
import logo from '../assets/Logo_Discord256.png';
import { NavLink } from 'react-router-dom';

function Header() {
    return (
        <div className="Header bg-green-500 text-white rounded-lg p-4 m-4 flex flex-row justify-between">
            <div className="logo flex flex-row justify-center items-center">
                <img src={logo} alt="Logo" className="logo w-10 h-10" />
                <h1 className="ml-2 text-3xl font-bold">Cali Race App</h1>
            </div>
            <div className="links flex flex-row justify-between">
                <NavLink to="/" className="button mx-2 bg-green-600 py-1 px-2 rounded-lg">Accueil</NavLink>
                <NavLink to="/racers" className="button mx-2 bg-green-600 py-1 px-2 rounded-lg">Liste des coureurs</NavLink>
                <NavLink to="/edit" className="button mx-2 bg-green-600 py-1 px-2 rounded-lg">Ajouter un coureur</NavLink>
            </div>
        </div>
    );
}

export default Header;
