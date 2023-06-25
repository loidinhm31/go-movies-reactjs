import { render, screen, waitFor } from "@testing-library/react";
import DoughnutChart from "@/components/Chart/DoughnutChart";
import LineChart from "@/components/Chart/LineChart";
import userEvent from "@testing-library/user-event";
import AreaChart from "@/components/Chart/AreaChart";
import PaymentBarChart from "@/components/Chart/PaymentBarChart";

describe("Doughnut Chart", () => {
  test("renders the component with correct data", async () => {
    const res = render(<DoughnutChart movieType="MOVIE" />);

    expect(screen.getByRole("progressbar")).toBeInTheDocument();

    await waitFor(() => {
      expect(res).toBeTruthy();
    })
  });
})

describe("Line Chart", () => {
  test("nothing to do", async () => {
    render(<LineChart movieType="both" />);

    expect(screen.getByRole("button")).toBeInTheDocument();

  });

  test("nothing to do", async () => {
    render(<LineChart movieType="MOVIE" />);
    await waitFor(async () => {

      userEvent.click(screen.getByRole("button"));

      await waitFor(async () => {
        userEvent.click(screen.getByText("G1 - MOVIE"));

        await waitFor(() => {
          expect(screen.getByRole("img")).toBeInTheDocument();
        })
      })
    })
  });
})

describe("Area Chart", () => {
  test("renders the component with correct data", async () => {
    const res = render(<AreaChart movieType="MOVIE" />);

    expect(screen.getByRole("progressbar")).toBeInTheDocument();

    await waitFor(() => {
      expect(res).toBeTruthy();
    })
  });
})

describe("Payment Bar Chart", () => {
  test("renders the component with correct data", async () => {
    const res = render(<PaymentBarChart movieType="MOVIE" />);

    expect(screen.getByRole("progressbar")).toBeInTheDocument();

    await waitFor(() => {
      expect(res).toBeTruthy();
    })
  });
})

