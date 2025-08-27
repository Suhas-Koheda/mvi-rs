import React, { useEffect, useState } from 'react';

const Reservations = () => {
  const [reservations, setReservations] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('/reservations')
      .then(res => res.json())
      .then(data => {
        setReservations(data);
        setLoading(false);
      });
  }, []);

  if (loading) return <div>Loading reservations...</div>;

  return (
    <div style={{ maxWidth: 800, margin: '2rem auto' }}>
      <h2>Your Reservations</h2>
      <ul style={{ listStyle: 'none', padding: 0 }}>
        {reservations.map(r => (
          <li key={r.id} style={{ marginBottom: 16, border: '1px solid #ccc', padding: 16 }}>
            <div>Movie: {r.movieTitle}</div>
            <div>Showtime: {r.showTime}</div>
            <div>Seat: {r.seatNumber}</div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Reservations;
