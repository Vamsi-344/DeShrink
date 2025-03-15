import { useState } from 'react';

function App() {
  const [url, setUrl] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleUrlChange = (e) => {
    setUrl(e.target.value);
  };

  const handleKeyDown = async (e) => {
    if (e.key === 'Enter' && url) {
      try {
        setLoading(true);
        const response = await fetch('http://localhost:5173/api/shorten', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ url }),
        });

        const result = await response.json();

        if (response.ok) {
          alert('URL shortened successfully!');
        } else {
          setError(result.message || 'Something went wrong');
        }
      } catch (error) {
        setError('Network error. Please try again later.');
      } finally {
        setLoading(false);
      }
    }
  };

  return (
    <>
      <div className="flex flex-col justify-center items-center min-h-screen bg-gray-50">
        <h1 className="text-3xl font-bold text-purple-600 mb-6">URL Shortener</h1>

        <div className="w-full max-w-md bg-white p-6 rounded-lg shadow-md">
          <label htmlFor="website" className="block text-lg font-medium text-gray-700 mb-2">Website URL</label>
          <input
            type="url"
            id="website"
            className="w-full p-3 border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
            placeholder="https://www.example.com"
            value={url}
            onChange={handleUrlChange}
            onKeyDown={handleKeyDown}
            required
          />
          
          {error && <p className="text-red-600 mt-2">{error}</p>}

          <div className="flex justify-center mt-4">
            <button
              className={`text-white bg-purple-600 hover:bg-purple-700 focus:ring-2 focus:ring-purple-500 font-medium rounded-lg text-sm px-5 py-2.5 text-center disabled:bg-gray-400`}
              onClick={handleKeyDown}
              disabled={loading || !url}
            >
              {loading ? 'Shortening...' : 'Submit'}
            </button>
          </div>
        </div>
      </div>
    </>
  );
}

export default App;
