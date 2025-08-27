import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

const MovieDetails = () => {
  const { id } = useParams();
  const [movie, setMovie] = useState(null);
  const [showtimes, setShowtimes] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch(`/movies/${id}`)
      .then(res => res.json())
      .then(data => {
        setMovie(data);
        setLoading(false);
      });
    fetch(`/showtimes?movieId=${id}`)
      .then(res => res.json())
      .then(data => setShowtimes(data));
  }, [id]);

  if (loading) return <div>Loading movie...</div>;
  if (!movie) return <div>Movie not found.</div>;

  return (
    <div style={{ maxWidth: 600, margin: '2rem auto' }}>
      <h2>{movie.title}</h2>
      <p>{movie.description}</p>
      <h3>Showtimes</h3>
      <ul>
        {showtimes.map(st => (
          <li key={st.id}>{st.time} - Hall {st.hallId}</li>
        ))}
      </ul>
    </div>
  );
};

export default MovieDetails;
