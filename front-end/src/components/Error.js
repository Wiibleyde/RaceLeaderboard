import React from 'react';

function Error() {
    return (
        <div className="Error">
            <h1>Erreur 404</h1>
            <p>La page demandée n'existe pas</p>
            <div>
                Vous pouvez, <a href="/">retourner à l'accueil</a>
            </div>
        </div>
    );
}

export default Error;
