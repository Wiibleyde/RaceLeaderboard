import React, { useState, useEffect } from 'react';

function EditRacer({ match }) { 
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

    const [driverFName, setDriverFName] = useState('');
    const [driverLName, setDriverLName] = useState('');
    const [driverNumber, setDriverNumber] = useState('');
    const [driverTeam, setDriverTeam] = useState('');
    const [driverCountry, setDriverCountry] = useState('');
    const [driverAge, setDriverAge] = useState('');
    const [driverImage, setDriverImage] = useState('');

    const updateDriver = async (e) => {
        e.preventDefault();
        const response = await fetch(`http://${link}/updateDriver`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                Id: match.params.id,
                DriverFName: driverFName,
                DriverLName: driverLName,
                DriverNumber: driverNumber,
                DriverTeam: driverTeam,
                DriverCountry: driverCountry,
                DriverAge: driverAge,
                DriverImage: driverImage
            }),
        });
        const data = await response.json();
        console.log(data);
    };

    return
}

export default EditRacer;