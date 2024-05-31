import store from "../store";
import { createUser } from "../store/slices/auth";

export default function SignUp() {
  const handleSubmit: React.FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const name = formData.get("name") as string;
    const password = formData.get("password") as string;
    store.dispatch(createUser({ name, password }));
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="name" name="name" placeholder="Name" />
      <input type="password" name="password" placeholder="Password" />
      <button type="submit">Sign Up</button>
    </form>
  );
}
