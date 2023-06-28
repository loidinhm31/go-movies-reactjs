import { render, waitFor, screen } from "@testing-library/react";
import MovieCheckout from "@/components/Payment/MovieCheckout";
import EpisodeCheckout from "@/components/Payment/EpisodeCheckout";
import CheckoutAuthorize from "@/pages/checkout/verify";
import { useRouter } from "next/router";


describe("Movie Checkout", () => {
  test("renders movie without price", async () => {
    render(<MovieCheckout refId={2} type="MOVIE" />);

    await waitFor(() => {
      expect(screen.getByText("Test movie 2")).toBeInTheDocument();

      expect(screen.getByText("Cannot Pay for this movie")).toBeInTheDocument();
    });
  });

  test("renders movie checkout was paid", async () => {
    render(<MovieCheckout refId={4} type="MOVIE" />);

    // Wait for the movie details to be loaded
    await waitFor(async () => {
      expect(screen.getByText("Test movie 4")).toBeInTheDocument();

      expect(screen.getByText("Price: 59 USD")).toBeInTheDocument();
      expect(screen.getByRole("button", { name: "You bought this movie" })).toBeInTheDocument();
    });
  });

  test("renders movie checkout was not paid", async () => {
    render(<MovieCheckout refId={3} type="MOVIE" />);

    // Wait for the movie details to be loaded
    await waitFor(() => {
      expect(screen.getByText("Test movie 3")).toBeInTheDocument();

      expect(screen.getByText("Price: 59 USD")).toBeInTheDocument();

      expect(screen.getByRole("button", { name: "Pay now" })).toBeInTheDocument();

    });
  });

});

describe("Episode Checkout", () => {
  test("renders episode checkout component", async () => {
    render(<EpisodeCheckout refId={3} type="TV" />);

    await waitFor(async () => {
      expect(screen.getByText("Test movie 1")).toBeInTheDocument();
      expect(screen.getByText("Season 1 - E3: Test episode 3")).toBeInTheDocument();
      expect(screen.getByText("Price: 57 USD")).toBeInTheDocument();

      expect(screen.getByRole("button", { name: "Pay now" })).toBeInTheDocument();
    });

  });
});

jest.mock("next/router", () => ({
  ...(jest.requireActual("next/router") as object),
  useRouter: jest.fn()
}));

describe("Verification", () => {
  test("renders verifying payment message", async () => {
    const mockReplace = jest.fn();

    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {
        providerPaymentId: "123",
        type: "MOVIE",
        movieId: "1",
        episodeId: "3"
      },
      replace: mockReplace
    }));

    render(<CheckoutAuthorize />);

    expect(screen.getByRole("button", { name: "Verifying Your Payment..." })).toBeInTheDocument();

    await waitFor(() => {
      expect(mockReplace).toHaveBeenCalledWith("/checkout/completion?type=MOVIE&movieId=1");
    });
  });

  test("renders error message when payment verification fails", async () => {
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {
        providerPaymentId: "234",
        type: "MOVIE",
        movieId: "1",
        episodeId: "3"
      },
      replace: jest.fn()
    }));

    render(<CheckoutAuthorize />);
    expect(screen.getByRole("button", { name: "Verifying Your Payment..." })).toBeInTheDocument();

    await waitFor(() => {
      expect(screen.getByText("Cannot verify your payment")).toBeInTheDocument();
    })
  });
});