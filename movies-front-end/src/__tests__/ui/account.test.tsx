import { useSession } from "next-auth/react";
import { render, screen } from "@testing-library/react";
import Account from "@/pages/account";

jest.mock("next-auth/react", () => ({
  useSession: jest.fn(),
}));

describe("Account", () => {
  test("renders account details when session exists", () => {
    (useSession as jest.Mock).mockReturnValue({
      data: {
        user: {
          id: "testuser",
          name: "John Doe",
          role: "general",
          email: "john@example.com",
        },
      },
    });

    render(<Account />);

    expect(screen.getByText("ID")).toBeInTheDocument();
    expect(screen.getByText("testuser")).toBeInTheDocument();

    expect(screen.getByText("Username")).toBeInTheDocument();
    expect(screen.getByText("John Doe")).toBeInTheDocument();

    expect(screen.getByText("Role")).toBeInTheDocument();
    expect(screen.getByText("general")).toBeInTheDocument();

    expect(screen.getByText("Email")).toBeInTheDocument();
    expect(screen.getByText("john@example.com")).toBeInTheDocument();

  });

  test("renders placeholder when session does not exist", () => {
    (useSession as jest.Mock).mockReturnValue({
      data: null,
    });

    render(<Account />);

    const placeholder = screen.queryByText("Your Account");

    expect(placeholder).not.toBeInTheDocument();
  });
});