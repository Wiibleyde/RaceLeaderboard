import React, { useState, useEffect } from 'react';
import useInterval from '../useInterval';
import { NavLink } from 'react-router-dom';
import axios from 'axios';

const deleteDriver = async (id) => {
    let link = 'localhost:3001';
    if (process.env.NODE_ENV === 'STARTED_BY_DOCKER') {
        link = 'back-end';
    }
    try {
        axios.post(`http://${link}/deleteDriver`, {
            DriverId: id
        })
    } catch (error) {
        console.log(error);
    }
}

function ListRacer() {
    const [racer, setRacer] = useState([]);
    let link = 'localhost:3001';
    if (process.env.NODE_ENV === 'STARTED_BY_DOCKER') {
        link = 'back-end';
    }

    const fetchData = async () => {
        try {
            const response = await fetch(`http://${link}/listDriver`);
            let data = await response.json();
            console.log(data);
            if (data != null) {
                data.sort((a, b) => (a.DriverLName > b.DriverLName) ? 1 : -1);
            } else {
                data = [{ Id: 0, DriverFName: "Aucun", DriverLName: "coureur", DriverNumber: "0"}];
            }
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
            <h1 className="title text-3xl">Liste des pilotes</h1>
            {/* If no racers show text */}
            {racer.length === 0 && <p className="text-lg">Aucun pilote n'a été ajouté</p>}
            {/* If racers show racers */}
            <div className="racers grid lg:grid-cols-4 sm:grid-cols-1">
                {racer.map((racer) => (
                    <div key={racer.Id} className="racer bg-green-500 text-white rounded-lg p-4 m-4 flex flex-row justify-between">
                        <div className="grid grid-cols-1">
                            <div className="flex flex-row justify-between items-center">
                                <p className="name text-lg">{racer.DriverFName} {racer.DriverLName}</p>
                                <p className="number text-lg">{"555-" + racer.DriverNumber}</p>
                            </div>
                            <div className="buttons grid grid-cols-1">
                                <NavLink key={racer.Id} to={`/edit/${racer.Id}`} className="button mx-2 bg-green-600 py-1 px-2 rounded-lg">Modifier</NavLink>
                                <button key={racer.Id} onClick={deleteDriver(racer.Id)} className="button mx-2 bg-green-600 py-1 px-2 rounded-lg">Supprimer</button>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default ListRacer;
