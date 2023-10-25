import React from 'react';

function Home() {
    return (
        <div className="App">
            <h1>Bienvenue sur le site de gestion de course de Cali</h1>
            <div className="buttons">
                <a href="/racers" className="button">Liste des coureurs</a>
                <a href="/edit" className="button">Ajouter un coureur</a>
            </div>
        </div>
    );
}

export default Home;
