import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

const Movies = () => {
  const [movies, setMovies] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('/movies')
      .then(res => res.json())
      .then(data => {
        setMovies(data);
        setLoading(false);
      });
  }, []);

  if (loading) return <div>Loading movies...</div>;

  return (
    <div style={{ maxWidth: 800, margin: '2rem auto' }}>
      <h2>Movies</h2>
      <ul style={{ listStyle: 'none', padding: 0 }}>
        {movies.map(movie => (
          <li key={movie.id} style={{ marginBottom: 16, border: '1px solid #ccc', padding: 16 }}>
            <h3>{movie.title}</h3>
            <p>{movie.description}</p>
            <Link to={`/movies/${movie.id}`}>View Details</Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Movies;
