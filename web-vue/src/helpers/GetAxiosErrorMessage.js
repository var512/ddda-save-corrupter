const GetAxiosErrorMessage = (err) => {
  console.log(err);

  if (err.response) {
    return err.response.data;
  }

  if (err.message) {
    return err.message;
  }

  return 'Error: invalid request';
};

export default GetAxiosErrorMessage;
