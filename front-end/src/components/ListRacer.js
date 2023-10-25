import React, { useState, useEffect } from 'react';
import useInterval from '../useInterval';

function ListRacer() {
    const [racer, setRacer] = useState([]);
    let link = 'localhost:3001';
    if (process.env.NODE_ENV === 'STARTED_BY_DOCKER') {
        link = 'back-end';
    }

    const fetchData = async () => {
        try {
            const response = await fetch(`http://${link}/listDriver`);
            const data = await response.json();
            setRacer(data);
        } catch (error) {
            console.log(error);
        }
    };

    useEffect(() => {
        fetchData();
    }, []);

    useInterval(() => {
        fetchData();
    }, 1000);

    return (
        <div className="ListRacer">
            <h1>Liste des pilotes</h1>
            {racer.map((racer) => (
                <div key={racer.Id} className="racer" onclick="window.location.href = '/edit/' + racer.Id">
                    <p>{racer.DriverFName} {racer.DriverLName}</p>
                    <p>{"555-" + racer.DriverNumber}</p>
                </div>
            ))}
        </div>
    );
}

export default ListRacer;
