import SearchUsersOIDC from "@/components/Search/SearchUser/SearchUsersOIDC";
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import SearchUsers from "@/components/Search/SearchUser/SearchUsers";

describe("Search Users from OIDC", () => {
  test("performs user search", async () => {
    render(<SearchUsersOIDC setNotifyState={jest.fn()} wasUpdated={false} setWasUpdated={jest.fn()} />);

    const searchButton = screen.getByRole("button", { name: /search/i });
    expect(searchButton).toBeInTheDocument();

    const usernameInput = screen.getByLabelText("OIDC Username");

    userEvent.click(usernameInput);
    userEvent.type(usernameInput, "existingUser");
    userEvent.click(screen.getByRole("button", { name: "Search" }));

    await waitFor(() => {
      expect(screen.getByDisplayValue("existingUser@example.com")).toBeInTheDocument();
    });
  });

  test("adds OIDC user with unselected role", async () => {
    const mockSetNotifyState = jest.fn();
    render(<SearchUsersOIDC setNotifyState={mockSetNotifyState} wasUpdated={false} setWasUpdated={jest.fn} />);
    const usernameInput = screen.getByLabelText("OIDC Username");

    userEvent.click(usernameInput);
    userEvent.type(usernameInput, "existingUser");
    userEvent.click(screen.getByRole("button", { name: "Search" }));

    // Wait for the server response
    await waitFor(() => {
      expect(screen.getByDisplayValue("existingUser@example.com")).toBeInTheDocument();

      expect(screen.getByTestId("add-oidc")).toBeInTheDocument();

    });

    userEvent.click(screen.getByTestId("add-oidc"));

    expect(mockSetNotifyState).toBeCalledWith({
      open: true,
      message: "OIDC User need to set role",
      vertical: "bottom",
      horizontal: "center",
      severity: "warning"
    });
  });

  test("adds OIDC user with selected role", async () => {
    const mockWasUpdated = jest.fn();
    const mockSetNotifyState = jest.fn();

    render(<SearchUsersOIDC setNotifyState={mockSetNotifyState} wasUpdated={false} setWasUpdated={mockWasUpdated} />);
    const usernameInput = screen.getByLabelText("OIDC Username");

    userEvent.click(usernameInput);
    userEvent.type(usernameInput, "existingUser");
    userEvent.click(screen.getByRole("button", { name: "Search" }));

    // Wait for the server response
    await waitFor(() => {
      expect(screen.getByDisplayValue("existingUser@example.com")).toBeInTheDocument();

      expect(screen.getByTestId("add-oidc")).toBeInTheDocument();
    });

    userEvent.click(screen.getByRole("radio", { name: "admin" }));

    userEvent.click(screen.getByTestId("add-oidc"));

    await waitFor(() => {
      expect(mockSetNotifyState).toBeCalledWith({
        open: true,
        message: "OIDC User added",
        vertical: "top",
        horizontal: "right",
        severity: "success"
      });

      expect(mockWasUpdated).toBeCalledWith(true);

      expect(usernameInput).toHaveValue("");
    });
  });

});

describe("Search Users", () => {
  test("performs user checkbox is new", async () => {
    render(<SearchUsers setNotifyState={jest.fn()} wasUpdated={false} setWasUpdated={jest.fn()} />);
    const isNewCheckbox = screen.getByTestId("is-new");

    expect(screen.getByLabelText("Keyword")).toBeInTheDocument();

    userEvent.click(isNewCheckbox);

    await waitFor(() => {
      expect(screen.getByRole("table")).toBeInTheDocument();
    })
  });

  test("performs user search", async () => {
    render(<SearchUsers setNotifyState={jest.fn()} wasUpdated={false} setWasUpdated={jest.fn()} />);
    const keywordInput = screen.getByLabelText("Keyword");

    userEvent.click(keywordInput);
    userEvent.type(keywordInput, "test1{enter}")

    await waitFor(() => {
      expect(screen.getByRole("table")).toBeInTheDocument();
    })
  });
});