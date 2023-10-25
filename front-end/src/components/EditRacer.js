import axios from 'axios';
import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

function EditRacer() {
    const { id } = useParams();
    const [driverFName, setDriverFName] = useState('');
    const [driverLName, setDriverLName] = useState('');
    const [driverNumber, setDriverNumber] = useState('');
    const [racer, setRacer] = useState([]);

    let link = 'localhost:3001';
    if (process.env.NODE_ENV === 'STARTED_BY_DOCKER') {
        link = 'back-end';
    }

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(`http://${link}/getDriver?id=${id}`);
                const data = await response.json();
                setRacer(data);
            } catch (error) {
                console.log(error);
            }
        };
        fetchData();
    }
        , []);

    useEffect(() => {
        if (racer.length > 0) {
            setDriverFName(racer[0].DriverFName);
            setDriverLName(racer[0].DriverLName);
            setDriverNumber(racer[0].DriverNumber);
        }
    }
        , [racer]);

    const handleDriverFNameChange = (event) => {
        setDriverFName(event.target.value);
    };

    const handleDriverLNameChange = (event) => {
        setDriverLName(event.target.value);
    };

    const handleDriverNumberChange = (event) => {
        setDriverNumber(event.target.value);
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            axios.post(`http://${link}/updateDriver`, {
                DriverId: id,
                DriverFName: driverFName,
                DriverLName: driverLName,
                DriverNumber: driverNumber
            })
                .then((response) => {
                    console.log(response);
                }, (error) => {
                    console.log(error);
                });
        } catch (error) {
            console.log("Error:", error);
        }
        // Redirect to the list of racers
        window.location.href = "/racers";
    };


    return (
        <div className="EditRacer">
            <h1 className="title text-3xl">Modifier un pilote</h1>
            <form className="form flex flex-col justify-between items-center p-2 m-2 bg-green-500 rounded-lg" onSubmit={handleSubmit}>
                <label className="label text-lg" htmlFor="driverFName">Prénom</label>
                <input className="input text-lg" type="text" id="driverFName" name="driverFName" value={driverFName} onChange={handleDriverFNameChange} />
                <label className="label text-lg" htmlFor="driverLName">Nom</label>
                <input className="input text-lg" type="text" id="driverLName" name="driverLName" value={driverLName} onChange={handleDriverLNameChange} />
                <label className="label text-lg" htmlFor="driverNumber">Numéro</label>
                <input className="input text-lg" type="text" id="driverNumber" name="driverNumber" value={driverNumber} onChange={handleDriverNumberChange} />
                <button className="submit-button text-lg" type="submit">Submit</button>
            </form>
        </div>
    );
}

export default EditRacer;