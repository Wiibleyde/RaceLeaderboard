import React from 'react';

function Error() {
    return (
        <div className="Error bg-green-500 text-white rounded-lg p-4 m-4 flex flex-col justify-center items-center">
            <h1 className="title text-3xl">Erreur 404</h1>
            <p className="text-2xl">La page demandée n'existe pas</p>
            <div className="links flex flex-row justify-between">
                Vous pouvez, <a href="/">retourner à l'accueil</a>
            </div>
        </div>
    );
}

export default Error;
