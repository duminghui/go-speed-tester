package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"go-speed-tester/rpchelp"
)

var accountStrSlice = []string{
	"5eBZm2xFsHB2pKzW5sV5JAwFG2ej8VRunL1rzRPTq9tE", "5JxRnnRwqBRdwK2VHvTcRLjBRhKTKSmZRVQLtVMyd17r", "AAh1c3iwbu4ns2YvxBRkE58VkDEu1vt7MdPxQKQqTSa4", "2wKENKmVTjZB4n7ZbFP4AZ6jtibvyQpqtXyFPeDbJ7t5", "7qYTPrrh5o2yiyGERQ6hKktey16Z18fe5Eei5ymhGdNb", "8SBrLHd68sQJgYGHipfQBdvfR72VqvXozBwEdvKoYWRh", "GAcWo7KvMqxtnFSULCCKkkhWAGnWZMaj9reW3mZqwY6M", "Aqj4KNN28TvvyGtDTrfwnjLb8eK1GJvM9tja5m9qdgDM", "HRk9CMrpq7Jn9sh7mzxE8CChHG8dneX9p475QKz4Fsfc", "2VXhUYA8r9dbajVrpYPhph2n4LMTHvPq9FZxePLojMh1", "FEXUvj7a46shp9PJryKBtvJBgGMbZF59uMGh4JWyEeK6", "7i5XcDRXEk2Fk4d17h7PCQ9LQDHKQuXCRVgLwTQjxkvU", "3fFqrYxDVaXh8MpHrCwnADpEf4tyVaAXFAoAhqwRc1RD", "38PRxtjdqPyDW21ta3PLhSWraUr7gH5b3wxic7HMURem", "murSwAXnzCDfDhvSN2qwmPdP5AynpQwM1XmrGSXxRV6", "GwDeKNzQGLimDszBhpikJ85Kzngpsor77ts8Ry2SEwtg", "74oaznphjUcFPQoE19wP9BoAQFrWctCjKiWjXf54XV9M", "HNDoGxAMwzyPCPD9tpUFkd4Uu4FdftuwTBNM73ZaJfsF", "BKh5Tsoov1dfoSRcnqsEFxRbCexmThNBWEMkrFr5SpVV", "DfbXsvDSr7S8QiVb9AAkzByWumLCPDhdsC6XmUfkevzG", "AWVazjpykJwn35oPBv8NYJD9JrstivUG4xqPaSeN6eyf", "6aEC2p6dnjHj3ADiuPg6tttsRXVsjE3os1owadXTPX8c", "HgovbkvirqK2x5K48UsrxvVTTSBUKBDN4ts7A3PoF3tL", "J7Gn2eT6QZihFxdREUGrixCNJkzXEarsZypZPWgTjAYh", "CXJrKKMcS1egXMPBTeL4yZY3BLqyuPUfN7HfSLhUxCDx", "CrWbfKwyAaUfYctXWF9iaDUP4AH5t6k6bbaWnXBL8nHm", "DNjCvHVPyfXdVjrYjJk6eUqT8jekk8XTxFYqiGEk5cXT", "GzSGyNXZpDEtvBMv4iBso4xKQs8wc8VMSj1o7S7WWzv1", "9SHYxRLjHTuksNXmRMdxvEzML6Vz5soVKSr7GaMk9ziF", "HuCtiVgMb66ypCE3cfkyLvXtUVmGMsR63inyWPbSZ8q7", "3obyhwwRXNRi2J389fKrhfyPuxht4rBgVf7hYG8cmwLd", "Bfsv5ATH7bTqp2QNguLx8fxkdDkPrbem7q9iYmARJHFX", "Fz8p9LA88pw3krMHMjBD3vTbQJjtPwAGNgwFAKoeKs5E", "ArShCAggk212K6s4gKEV6J4RM7i3oWwRh7wdLctuXTnd", "z2KxiSejQmNNsyxLFHbrewNLDeGLFZahFNSLYht2FFs", "9ZVAXWGJTzyvGwKEwoiE7J3ggQT1X7bYiVGPoskf9N56", "DVBx9FyJ1JRzEi5oRfrZdnxddwGHF8GWvpqwwxwejb9P", "9bcwDk3mAq9NGwMx9F1hjTzV6Pycr3KRJrfGSb1m8bet", "5a7mrH5HCcaDT1TN1d8GsgbrJft4YGarMugq35pYte4d", "8An77vrD9NHLRiznBiBqoMhNRXcZC2fGVsdWsVjjHQua", "7hFDdgVwEDhw8JAkZ6mENJ7CjCYZZNgGxgrXAsrajBdG", "9L8m4LaCfTLoh9W2ohJSp3G6oNQa1wXb8zJAJjnbyavM", "FEFzBbbEK8yDigqyJPgJKMR5X1xZARC25QTCskvudjuK", "3kWwdn7WzysuBvCSuAPXEq4N489j4SMY9qC3QaZHAXmw", "4xkiPy8CyMA5HHE5kTEodUJpBNMqGasKeFWv252sad5c", "8Y7cJU2pqzvPDWzjdKM8dzM6LmgC42Q6DpVvsq2gU892", "57GkWQBfVeRcqp1yiSTj1i6AsdS3JtrSBfHpVbNpbE5a", "xyapap7izq3no9ZFDqhDAWiZVhKhYeHYzqRHHfn5Z9C", "3atuP5hSnUFLpH1JEyGcFRJ73cv2WRgfwTtSaHsZH3uf", "7YSGSuQFBZQh5rGoqSTwY4haLX4nvQGytutPv9owbry1", "2B2ZiwJqsMpiqMnLfUTLvsiQfZ7PKHLZtjtvZuwzBGQ1", "GBGxwY1eqBJcTVAjwFDpLGQGCv5eoQTciudT9ttFybqZ", "FHAqAqqdyZFaxUTCg19hH9pRfKKChwNekFrY428NVPtT", "4GN51ZfmF4wzXaGtReQcBB3KRtoBoFomgo9gUBnYkeh6", "FRw1XgTitV5J2H2Fz6zFAUAPku2gxnCfd5GVo8GfzCPU", "6b5tQh5GTf4Rt2VMMDqrr84EfXezUMPj3KTpVgebzbUt", "qXeo1zzAnqHseW8mPzx68yoro7Tiprnmg2Dkq61vm4Y", "ApCWKVCMn2AW77NFDSgynFP2fVT78m1i4vKrjdUci5Nb", "DQX9NhwznyWTYcTJ8uiqZP3PrzqRmfGNj4XNQzVKG8hW", "6Kh8oeMr4Lfjog6MNS7kXF3GPJViR68mFqc54dzjto5s", "GWGoJyxjZMr54jcxzpVi1ZBoU3zQhrDwkuP9THdhjvmz", "Frk48ntSNU7fjk5sexecvNi5DrsR3DgDcPvbbgK3sfE6", "CKeQZZqDyS2awcaCDrN21BacbM8WnraWEwE8h29siJxH", "AzuKQFSUxiWhJhafj8CsGtcWbq54cBLDsTQMxNqdRHXj", "7BYSnBn5GfG9tsJ6j1dxRtuf2B9zmzr4XkkxNdakNgaN", "HRXGvc6A6fDwNa8M2CKy9FnWkXqsABXip36nG5csyboi", "7z6ZgwNnQ2qTZasiT6vvN8snANo6DjSwQPsRSAzY2GVB", "GBT57JSwpXkicnpJSxdh6uodo3dF8zCKURqvrFnKPGAx", "4G9Hp77tNNfMuYgD2DfFxZAAaruXuJigfbU2FjBreSdn", "93839nPc6m7AU1d8p9FHYKVdQ8PenpuSn7mEeXAPShFK", "CBdb31UyA28bSD76ujp1R7Csskxs6fLrgHYQEUaFWZuX", "7CrTDvxmtJ4ZecuHEKbgUnV7t5WXGr7m6S1M8CPfGd7R", "4eq3AM1bL9WoA4xFwPLeXRCGAFHLrkQXiNuuFhxx5y4Y", "AeHQKV1QA9TqrF5yu3zALNaLs51r3iRsq3BkL4z1WMnT", "9D6JfNjyi6dXBYGErxmXmezkauPJdHW4KjMr2RGyD86Y", "HsEKAUogQej9FdRRkBNnt1FfCkTzgfhN6iVsDpCkAfdv", "DBXqDuHCPH5gq2wzCv9MAfxXjKj2qKo71Nb6WWrXEM6T", "J9vQmVKt6tn7aFFRREykZYu25nbnyi8hDoecLdBTow7", "6fp5FjuJipRbeAALnMyfNx2vrZWj9ywx6UhAiS5dNQFo", "HDQEebBRfQe5dKy694ePaNB6wbQqBpj4UX4cEp4xUo5y", "5B2Jtdbj8LKvv6JKMrdJ7qo6GBLVVXuXc2TaQdv7dsPs", "6DX194ZeSQ3627aABmcuC6tWUb2UFhVXnxaZKVMsZkFQ", "HaynHdMQd77Wph8vj9vsjeAV1Ci9URg97zpR8JxZ8JwJ", "DQyrAcCrDXQ7NeoqGgDCZwBvWDcYmFCjSb9JtteuvPpz", "GrJfBKuwvmoCnTfiXSzMi9Ga9Ru1HEGoHDCgyu62Jeoc", "8Bk7PFJrGuSNX3RvTyFt3cF95RHm5r5yvHHygyuZJy2V", "HTz7HEEWN6kn3dSm6nrjQPJcXxjrVNWkxMDGULAXXrBV", "4TwfdKtNp8L2cQjPuyV3gebGw6yrMi5ojkJyodL3JuGW", "TvsyDByZiNoUnevxqz22F37yjkQeKTY3jKXxjiC3Y5t", "AseLV5kWbAjNETCKJsXcrrs6ksvBefEPdRa7pKXFsvYE", "Cj4EkCBXSy1EsgaKnm5jibYHk6VcMc4zz93DVinLHnPA", "B95W2PrRrEwHXnv3SiS4SvBg5W84maUjsy8omZpbAxW8", "DJ4jXW1VJ6rU2PEGmoa7BQiGPkYmK3cT1LHUSiPy6WTD", "DTyoukHrczFNp5U89pdmNMSTVvFMLY7gemHoYksqmcrr", "FQnumCbTodhzTRun2ECc3Xi3bcEh6m68EWhJYFQMZase", "Gm89eHx6PgrMC9Nx8yWF4GVi396Da9bdahhZLG7WzAQT", "FRqK65KZiyikNZRJD4H9PtgCk4iVFeHfTEjUxw9Yok6F", "AZPsv6tY1HQjmeps2sMje5ysNtPKsfbtxj5Qw3jcya1a", "2txPNVtwCHnWPwcXN1yCrP2RTDiZD4A7UA8xsUm4Crq6", "13VrwVRxnsoMPHcvtbhS4SgK6s8s4XhZsC6fnbK1ythG", "8KthUTcGkhn9AhnwjGk38Ww166mjWi4Xabfv2FmNzpZ6", "GzRuDJgWRJkqstUAjnRxZeos8ZDvoZLuCCuzYfeXqZDC", "ArE3Xor5x2EuwBFmyHdonvHMWeT7g4vGAqgTLWamPQkj", "DWrx9U6hvpz8L828eMP7tvZJELDSrKVNdU9Kg21VrZYf", "6D8Lufu7QRs7PcEn6kGKf6bvRDpiNRJo4KEZPmr6AQfo", "2LWp1Hjg4zoba1pd96TARmQi76aFnjFTduoc8PzitxCB", "FdmKUE4UMiJYFK5ogCngHzShuVKrFXBamPWcewDr31th", "FkcQmYjvfov5T3rSmpSsMAXnDYZcP1A1uY1Xnu18hdwo", "Akrj4BY2JHLExvKFLp3NNmr3P42dyQSSqsenj2Pgeug2", "DmV5kVwtGHqpgMzNcLCshzVkpsZTbhWyLRyACvJSgnYg", "ssnPNyhAXz8bMXXfkCwgEzCW7dMzUD6r5r379S5nCGU", "yDd4FPYvh5ESYwTRLhg1povJ6jiyMrfCd9GgSghZx1c", "4CNkiWouhXwbX41ePkbrp1xaXRTrfvhnEshS5ac31XBH", "GMfWwPEhV6Xn45SZPFvvib1HpDq6Ni31hPi7BQnp3Ugs", "DvVijge9HpEpfNVk8cTzdq2GuuF1eUcpVPGLf871HXFz", "FSrZByXs4q5xUYgkNvfk31Ecp7p9tp4r1ChRszH8Ctoj", "FGhnuYBenQZmahgCkPvefcwEjyJP6SxY7HrLDGGecN61", "DpvcE1swD9368bVzNG3RfbibfjNrsJhv5S2VTxdw2YCB", "DkDRaPfVfkg5ocMYFX6HyDk9xjia8FmKp6m6puPH6EG6", "Cgr7HzhB56E84eBku3jQciJZSNATrwribLs1p2hFcnA", "3NtrQ5y5XpshkJwDMCu6sDwNPckxpVYLBjePiUCMW6e8", "52DTEuGxBoMqCQAe4Qc4EgcKtbXqiSPYTeHky4p2QLox", "E8UySVsiHkgcec89P9HJYST7d1ChU68hB73nJKp2dqmN", "CVioXLp58QsN9Xsf8JkAcadRmC1vsW74imLpKhMxPWSM", "5CUeyhmQtqREtgmercEYf6GoqSrm5DLCBomoaJ2HvA1U", "FxZxntjHyxnT7wy17odqvSgCWMS5WSwo8LnkewuW1K7z", "DM528LQu5UxZAiMhZwaTriPMGggG6HTjJAt4Tpn7xbHq", "CdX66ToPzgQPGLtUNx6CwduqqdbZe1FncpHvDeRa4wDo", "3WcfMiQudsAPPgYGweHiS1aLyHy6q4uFQgbvEuMst2xL", "9BnkuYqwYdrp7A8kV8V3uPRb21Kjihc7C8eg6jmMz1dm", "5HohDejrRRArgrnnbX75giZ6ozfkEa36AGQgnoQZYAwx", "7QNuAgGZUDu7mjh3cCEoPywmDY7QzXTd45xh7QGEpDhi", "6hNjJ48Me4uLhnubuvvsMStVLXrB1qXw8qi9u7kr8apf", "9bHmLQv6D5vQTpgE556th2u7Dh8YW7J7ZqcsJA6rUbDm", "2eV8JxpVaQFKdFaN63u7ZRm6Bi2jpmcLebmxQaw1Tcbh", "J2JvcFhYUbTuszvXAs6iEu8jYvATCdkK87s5qjcLcLST", "9vhNmymWtaxBtRmHvA8SqAKcTCCPtuWZmoQXPfQcLvb", "9VwdsbBGkB7V2aBaGwcArKciuQUiM2ARtSZPSfTo3x2P", "3RBfKkZdeFjCuyrzk91uKWEGfqu4eftjo4UAz1H8ZVNE", "38vmV9F6UnNEMqxDiRdw5fMMXm9dfKdYbNzYqRhULWUG", "2wkkDrjQ7B8DwGfnHkYvmZw5sovW4uHKzvLTHgBMf5BB", "2K2GkLihhVDBZKN4uy8tTKHGQHaws8Q5XvUhqRnZwAmU", "Fo9LbJqwQbjZQqaatypy3qM6wnbMddmd4nVYSF7xs2iD", "6sAEqKqqHAgZagEDEXQXLaDMDmk2KmY3B3KuDkwG1Dow", "4ZkH4CHGaAvXLV7PHoPAYzXqQ4Gt1NbtyVLfwowFwvbN", "ChbDnhhp9XocaZ5T4su7ZFN4YnJZZGTwX9e5C3UDM9M9", "BcNn5ympeqpu6Gn8vHHCttRazkjCk2Lz7fTEAM9hmSFX", "58oQChx4yWmvKdwLLZzBi4ChoCc2fqCUWBkwMihLYQo2", "BHHgeXxUE2ckMifWk2RQaTWKUDd5X7b35VhNoasjYsRW", "848GaUFCPtEBsJnFzRJwoSyMm5GQZcnTKQpEBM7G1wJh", "6cVw51aB4WjtnNGWAmD67DyJDGZqKf7kifufqUtntq9B", "HveQMxVRs1Zo3bcSbT63bd32cGuc3EgHjKdiHJSSMnf8", "37khuexnbYXhvfkjoTfq5V1P8fNzW1LBiqfwG8DwoPX2", "BY67qBmyZsvPzLo1cPayAQ6smCH1QZqHMEuQeqLAAesG", "4ggyAS5D1jELuUC2psCLs7dzAFUeZoJiDRVoFpCp6k5y", "38hDp1nKrgqxphg9nAtqShBubmT4dLRCy3WiVudDGhzQ", "D246EPLEgcd1PtZWNPia1GWo61muW9ZJ8LBEYynCAPfa", "Fmwt7CQNzMNfwYQKTvtYnGZ5M5ub1qjzyhamwRoXQWcD", "5JjG1Yr949RhCWL9kj8GzhqpZHMkZeSmm2EY789isfaJ", "BQu2mCsDcp9aRDCMJoi3ttVKX6rJ3oJ6h5y8sygB7Gjk", "9ocRA9NYH1ktQP2Xw5ikg6JCtDgZ7jujY5wcvU77xDPH", "CmqQmnD8AXaaADZvvKMfoCE9tPq8tELwSAZdmRnYs8Kt", "EjR24EwAyELwv3gHuL9BJgfUr2tAH5xToiceBfwAhwJG", "DmmSN7NH3FpKSkfNuE2MbbWPoW1uVrFdJDVtrCeuo5Wi", "9aoyrobD9PWPL9xVgTwnPTFfmVA5FeygpwCDN4C6Keow", "3bViVVcUtdeMgNBrNMkK8zTXrFrNZuqGBhh8EgiNrnJf", "9CyW29kpwZv63JgeDCPjwUZ5okE8J7xLj21k8yksgXL", "6UoY4kBxgYwgSBHW8Us4hKMVhjPr25JHSbPiijjtuvP6", "Da1jx7BFUeMxEPdKnGQ6aqRWMn7PAsXd9HS5jbMQGRh9", "FdiugGGKaoJeWVXLVGWrvashdtsAFJJUT7AeUSgmoe3e", "8JjDjehzuFDSDhY5KYbppArkW4D1mh6Wkir4PRRBdvvW", "5WsRMsS3ymvt1Zmp3H46agmi53WtQhyjSoPWuCRWWNyR", "UBFtqPv9QTTCAB1qTkGEKbZPP1MD5qjZshZVPgbe5b6", "GcM9qac6N1s9yZuNdRgHfaWa9UNmFN4j7ZESGJA1kxgb", "BzrVaaVEGev8nSfuFAmxqNYrxAnRvLvsBsySsiyiAYzU", "DAYtzvKXKATohUUaqj8hy86fnu92xP8aNQv6e6WkiTuU", "3rY1TU9NEkExmkVt9ouNL4FSoLdcCxmNp8rbQFaY8MQa", "5rsUTRPERNWejd4icM3y97WJdRiyiDFcNxU8nnFmAhhj", "HwuY4vrgc2mJr5zMHfQeQu35cHUf3Ngsy1dibWGHkYis", "HVLu3mkWyqaWfpbUQtMv2viCD7pg34JjZNBbU9LsLze", "3auxqCEdeP3pGenH9YbpHCit9uVRecWDxwnCvxxNxDfe", "6UmmUiYoBjSrhakAobJw8BvkmJtDVxaeBtbt7rxWo1mg", "7uusHUjtg9duoLZw5e1JhX3KxS1sF6Rg9CZ5Hs9MfH3R", "H1uVMacCaCPBswnBEnyMkZvdAEcodbvu8e1rv5v3o6a9", "EKVfTAkfQrSySTK77nv3TUm131dcdwhjKZWgqUca7KMS", "AdiSyFVyrNMT1FkuuazgFXPBWRWf5AHoncSgNkr5h1mc", "3oqommAn3DLRJ3Fx7Avc4xawHUKLWg87KZ7zmHGvgtfn", "EQMVw9zbCvhWmGMp55oA27v6pVVT1eYLNJsYeoF7beqj", "HLmqeL62xR1QoZ1HKKbXRrdN1p3phKpxRMb2VVopvBBz", "6BDCwrz9aXfywwBYTvRF7bPsu1JdVPdGpzvDEkRRk9k1", "3cqCs9bUPoXe4RjNJ548dHke65UWG1AidMbaMnRhGTZV", "Fgz8dRYXsp7Ntw31WwkPTWV5YSosTmM7pHgnjQyPFtgE", "35vZzSoEUBhUZnKYDvyaei5GQ83mP1bv2WcMdmhsAThH", "H5HBNgS1FyrGYWjXbjfQzYG2PpN4VcDDHxTj6XkcVxoj", "2VE9wZQLNGGXyjSyEUzgBvLxFFupp9DgWBUY3QhFKx5P", "FUBxj2kgkYhBmJ9mxSZC1BeSx8pX3nRdthgujVjJ4ktu", "7NDfhD9hEkQFfVXMTHEi3p4khDjBFjycc24Qn2coZwgr", "93Wny6vmaBSADjBFhscGLfK54rKbJq4MHefPcYXtVfvV", "7vdztMUduVQWvUcttZqpBguTtkDveteWt7oHHBEyt5PZ", "Fdh7esSmXbjod8h21ZHDhkbKfU1ucy8JXKCoQZfuGf8G", "CAjqhMToq86DJ9Amy5tVdTdpBB494NVxhEEu3ZPK6SZm", "5YjN2PJTqCU5XcJqsDqfAn8JoDyjwZG1t2TcBS6pgbqX", "4YvjMvJj3fNqYXf19e1MdosmU8jxGd4CESQRMpdS9bTH", "768jrd5iTinnenHMZedR89YXrGcEqK3fdMeEbgUBqPPa", "Ht7YKtj7Kki3pSxRmJydFwr6zhApvE2Vq6QCUABiC7Wb", "5TadnacfiAVoifWBAMbgYzenP1GnRDBmTK62bN39cwis", "E1rv3Niaz2LHUHFrhd7CTZeEhnmpWuBASqnNzzdL3hGN", "5zSMP1GvuEqnr4AaZNbBDn7zaGe2WRUdapzMHSHVZSEG", "F7GqehZ4VP1H6SyENQYDyddgxFSkEkbJNVCr3pcJdqJ9", "6s4pyCwZju6PYsvXPoVYdde4s2r2BWoj3kr8epQsmoPF", "Em5HLjRDwyJpUdicGTTV6rVyji9tpEXDceVf4pyu5vVF", "H1yz9yZA8d78XcFtnYAiAK5UBi6dVufW2oeTXZXdLt5G", "FHofFY7uiXY9MAReNKQ2aatDQSs8XMz6uykEfdRz43S6", "5M6zJezj6Q1LUcd3ur6rE1Gg9GzJ2eWVp7itTjNW2GZg", "4tAVJzmzMACrp91sCRJVRPkdmuCbAfE86cJddxcAAo7m", "5FmRKxpFuWTbsfUywMZMUiJYhvhZFYFG5c5dg1syFNrC", "99yyXppyQP6rwkw6VsRZ9hHHNCTWkmuTZFag4Mkx64sG", "qfq7Ayr5vBCukQHzyr3HryiJHMYUFSez9FJsezsUMmB", "Ac7oo3nPb5tbAzsvK6ydMMARy5c8C2MiNMyR7B4KFn2q", "CuWTdGwwGF2TaxuddQwB6yHFjQbDpLLE762yhUfVse5f", "9r3vud5trE2uRtUrjY5LVgaxcKXcpG7h7jnnTUuYpHq5", "EABExwUTkPhoeUcbReHJdx4UiMLhNaqdCAPAFwD1qZPk", "7vgAhQ3JmP3vkdAjQAUcyLev99DQPFaRU7qn9aYhGpPb", "BfioysFSPUAQhNpPedMgoT7k1nGbkaAju28YkmbdUV92", "9e2GPvARG57JTgME34HEMtRdLyZBPozS9hu2GiMKGYZs", "BqeMD525qVWbhrmaVGUZ3q3HYsefg8veYbN2j1wbLZEE", "EGKKv8nJ78VeWUdx4CXNisZdgKM3PC9CJCBfeSiJJ2sa", "FF9FPD4QDFonF5FKqSL62kj1NC9jikTwQdquVLcziy7U", "Eqrhxd7bDUCH3MepKmdVkgwazXRzY6iHhEoBpY7yAohk", "8FaTYA4mX3xXh68ePk4LnkdAtRZy1FJM1T4MSHbTxY5m", "9w7CqRD491maRWGoEdWTEyn86nmUhfzYxuLdD7xidkcS", "7rWoQ3WNU2XPCGVxsRo8TyW9uUFQonT4zCMaF5mqv5H7", "EgrHCYitzsCkNt5rHb4MWQ5WUaYm292NiPNjkqzKoRHE", "4M1exGfFG8AnVXn7xwScN9Jno49pfChpa7e9jF6k9QYK", "9Nahi6Kdh8BAPMjQ8hKqmDRSPUtzw5rb44Y4Sa3kRYMy", "EK1MRgnyMYsCFf6LCj3RdjMJekadwYkbTVfh8TzgDMn", "F7kjsUU65JtuKgTAkJFaXKtqFcjqH2bT6sND8WXEbpHQ", "4Vy8tZwkwjcdcmqyk5HgYpMkhsueUQw9knSVbhZipoq2", "6aQZi9Ug5PZxe6DpgzPYChXhAvpgGA7cY2J1che7ag4u", "8spM3hQmJ49hq8RimyHVBzBjNz47k8WxXJC9DdEzPhtj", "EoCKWrk2k1RnQYDaZKUJjpyy37yJpxJTEVJ5c9zXn4jL", "E9Z2JeEKS2WGGyA18mGU33rnQskK9moPhM4tdzrv24fh", "93if7merLqUTEKEKTAFDUu8Hsv2yRpcC6zsZpo3E8zuX", "F3MnZUMGihAiZHNk9UGY3jYWyjywPCmqBavmLep8TvPv", "AjF5GsH7KaQyhC3cTS3JN4fuczuftWXpnLWusSPnGbeF", "2BXS3XJMTFWgSLf8d12bKwNsyXuUX7RnSsRfYdtLYxoR", "GmDGxwD7oqj67LaBTAUsE59g4Wgr9E4ooYSSrMh72Tse", "AQ3rjwJVJcVetDFTLdbP3GWo2U6FHWmnWfHzk11RL6do", "8wQrJEcAKVShwHztQywjMowYMrneDZXGkj74j7x6h8Cf", "6HtYg9mC9ZCh3r8fggz1SobdLLgZMRAAjMvCFnvh1EPj", "Hu6y3hgYrSKaZbV81o2xEMdQGaoLWsfZkP7zHkDam9uD", "BqTsa5abxyJtr9hjQkBJycnKyKyUWGcYiif7z2aZh3at", "57im31Vy5Ve8BzQeHpxYBCZTokZYBuW1JvX4x9gBtyxd", "ZxVJ5M2PHCzGhZHdnoFRwmqHSSHNZ4ecznYeAcLZq4R", "Egrh4WEzvDUoGGk6ZfHYns5mwqFhQ26a1SbXPUMNumNo", "gng63EZXkDhK3Qp8KgvLEZkcWmVDrmBe3EuYRy8mBPy", "7D4wZWAQVJQJSgaBb6uHpXL3X6wjaoTtPYQEyzZyB3dm", "8PAAfUWoVsSotWUGrL6CJCT2sApMpE2hn8DGWXq4y9Gs", "4KC66wUybZdS7qrEXyoeGXGnyVJ9FTo4wnY8thH9bdyQ", "EJy2E3nrGhmM1PcC6SNG782j7ugaUqJNGoBaSXy7nZnR", "B7tLafxXDHDLpdxuvcFEcpYZ19iGWqFDZuV7SaPPrCsi", "HfBK19mBWh5D9VgnsPaKccfQaD79AYXetULtwLo62qxr", "EqiggpHL7WUxVbb6xBF3UDnDpEE4JzT5sHsmDc849rFa", "DwNodiwvYAXJxfyNVs3V3mzRkn5PdEbHTZKtomLe9Xfk", "CjDGfdD1WZGVbkYcXTKDVk693Xo2afQViFqimGQv3PRU", "J8u8nTHYtvudyqwLrXZboziN95LpaHFHpd97Jm5vtbkW", "BSjutQ4SNYoSaBKovSmmkMhf3v7i6h2CKrrNpJoC76C", "94goPSQw7uRm3rVphjMK8EmUcHdQ8qfvHkdj9bUiYM9N", "2DeoHHaKgX8HdvjEfootLyxJjqWdkYMsFvJwtLwTJmoq", "6HXfUDRXJkywFYvrKVgZMhnhvfqiU8T9pVYhJzyHEcmS", "3gHsdLrC6i8j5d5aHRJ8k5gKaq3D9SwWdWDpSW5sRAYw", "BtPwh5aUFQY6UNKZj2cpuUpD9a8JqioN61qhis8mqU6t", "HGNcDqvAskFVaVAEvR4Fi1zddJLNWfar9bsSAn868Yfa", "Hw5AvLXZX36ARPCzibr4gkyj9HTrjbbp95xrkyci1FfT", "KHcipLVD2Ap3yJzHbVGqgjj6zsP7h9oEHaRL6Ehxz1d", "64uBwyUUAQ3X4imAV9EKXUyb7PsPFDqN1BfgekEh3LAx", "4CwZg46Bpg8VBKtdKqaQ4UzyvoAfawocvgueC6ypq2Pk", "DK6Que2ZcWDARnCqYudVru8xVGcCYFGTcHe8hdZeP3uV", "8fbeXQvWamUixaVBHPp1BKBsLJpm87LAeRev4CGSMDJ3", "6R9vrY5QBDdudrNorf5VFRcWdMkChAtWcaAbGwHi8opw", "9J1wbqE1BHwKKxErQSMpVtCxBEXwrzc5VxmU1S82NTN4", "82a9AvCXqrx6rJuT1QfjuiZ9qeGWuij83tDbatY96BXd", "8ZHyUT8BYhWnZcGGn8QAT6yFSUBSBYKLgQDDKsNaXVfG", "G3CikvhQ3NwyZHjWjCF5A4QUZczebTY9SkLaWPNibsLM", "GUMBzhhoxtNi5KedB2X6ntBCiSbutX163qi6wZr3acnu", "6UNiUrpd95JCYDR45mQJ81h9wDDnQh1aN6ZNPiohKYPd", "CgQ5jb2ztEyM3rLY53mMBAJ4GJfkUJr3c9CVzqX6jZcn", "4s8CDoVsbMzF1vyPP7WwtPRsuC5qCvCvZgGrfxz4bJ1W", "FGpRgFyk3WTsvMfya7gww46QdA2PWfiggkrQWjcT3Wam", "FKxiXWR36aGR2iLLZRKfXuDcQvAHA1cteWWLEHtGSrNW", "6z71DBkkv2rQMaJpwquCdYLbSYDCpKFofLzMi9GCEX6S", "FoCdbPmPyxWRTKTN89qSZ2GNruFef3en8n7VUHYoGTo8", "CwPQ7GFk3rwjpAJSDftamkQzYH4VZbb4Bspzz4fTGTzJ", "GrjTEBKWxku18PtS8WUojwHRScAk46urfmgQjDb9k8o1", "44XiK3hmWDk6wV64yevJ4gKg4NnW9dTrHmKFZYPMdWTa", "7sPDo4E37dVwmFPU4Xba9YwzQHuEpYCsA9W4W9JkxBH8", "8NAPmMy5HXxKqT5r39uqzTVZ8Lw1E2h1h7E6oiGvaM7T", "4zuyAKT81y9mSSrjq8sN872zwgcD5ncQGyCXwRJDn6tC", "5YjfrxtcnaFtHWJUUNjh56Tb8QBCQSpUp5s3tEynZuNm", "9nL7B3trBbMadc3udUgtXVDWC5oHmVuLUA1HDNg4EcCR", "9i4H5SfpcYQEFiHFLzH851v9ry1i7GRbXpqL7AMnByLv", "FVDJiNFWcZajNUZs9ApZ5V4NinU47M1bnmtCMWXxN5LD", "4JAy22v17FGtbXindnU5yQ3vbnJSbeLL27JBrXT3zdJ6", "GJdfkP47DABLfD6jpV8wpE6Cnp1FscNddtqdwUikipzy", "FeqYC1Fpyon8PedGjApej9fRQUGcB94byBmsSmsifpRx", "DqvEhY6CpwhmF7ru4gL4CBdHEkC5G1M69Npy9PHwhvvJ", "CVAR7FQx3nM9qQB4ps3a82zhNNUVQu2Y1nBeMwQyBA8q", "H5AZY5uwRpN4KoUUNFzt4krymugd97TzL7QoW7ikA2vz", "GSfBWHRyckZtFzqYpVcL67JAbjnYY3WJhaCuz7NwTYxd", "FbB8rsKB3bXNC4GVQz85E7g9adJb74JbkVWciwUQ2enV", "9Vo1JLWRRaubF8YEpcLNcBovtrNkYnNcdXDBMNUhRbgd", "7Y21pUkgRaaTYZj7aX8oes2kJ9ticy1Zap1WgdfcUaBf", "9jgVdv6m9h8BShSRnkuyb6xQwWxfnsHCzpCNA4p7zcwq", "5WTSHBeNu2gked42LdyGaf3JLwd7aRrJzu94dhrDKyFz", "4ZbVzZgH35TDwDPemqKHv4B3m1rzobeq7HWSfzwTHtr1", "HkFSkhEQxiPLvP6iUqvRoUEhUKmm2qzbtcDiLCDPx95u", "AsQxZ59PrWEKv6eD5aVk4BgvkEQ7mdEf87Fh31Rr9uNS", "BayoEaHS6ssa2sivrXpUPrp252XXr3bocJZQjRzH29ew", "B37nLT6ExBw5r6BRK55BpGpBaZNskKJqVSET4CocQF2p", "FdwdTtLAPWZozttF61YgvKahdcpo5iE1pHWggj9zomqK", "8fEdArAuMR3b44WQK1UL1fDLB3kDK2N9whhDRP4sWP5v", "CRm1ooXR3VCSCYHQ4FaCx6HNm1Zs8j2EjhiSpyZyfL2C", "HxdzzXkT3USAVZmHfEum7t1aWL3dSj3JFcySxTB7sV7L", "2TcyA9hAw5jVhkYFqAkeNwi2yr4qfUytA7iTpfSE8Ewn", "2L5CC1tUjo14c4Et4rUV4LcubX21mBduNmC8WdcsHVtH", "EzN8mttBMhrAqPpjbjhVV2WxYwsvA4GLTf71VYyiqdUu", "AB4wFKn2rTJPRG7ZfR5p5c42dzSzfvBZU8ANyAcL8FXW", "8kDepzLGq8TaK2ATEePEBP5b1TXfhvVdqXaY18tRxwPj", "83wre7dTbbNKoV4E4HEVQBn3N7h3evWiuGLh3tg9PuWh", "H4zMatEWC1cgzpJd4Ckw29M7FD6h6gpVYMs8ATkVYsee", "Ei23wxsu7WVsXv72yaTohSVASLqseinqA7DqXktprSSz", "3RddYV1xRWzPDGcXysDUdWCEJpyWBMhoZx1NnTmY1zzD", "72zS4UiCZnnoXuLprd9AT36VDCtjjaBnbZig55e7agWC", "XoiXdzxuTuXFxxSS9RCzuHwA5DU2h4vAwAHZSg527R5", "JCu3baEcUyEkxFKb856eLfDiVf3D4xjDxZSdYNE91QSD", "EPqH3r4wg2KZsDigAJWAeViYskBZ5HWX1gHgHSU3jkWz", "DFSi6Faz7fX1YWfDRmB5GxZN6soekbyTS144FN5HPLLh", "JBSk4Gb2yUg763X88XfiKiXKuEwgSgGpDyoR9oYftwEi", "5EowrDg4UfPTWRpwo73QgZR5LPZ2oU8FRzRRMwf8mWqF", "CAaZxgN7vgT1SFCWr8dCj7N6QWkP3ERHSFFB6xPC8mup", "7jzwUCSq1R1QX72PKRDjZ4xgUm6Q6iiLW9BY8tnj8wkc", "DrbL99oPdR2N2bLcLm6gv3yWaaHPgv2vNAYXHeXHQ5U5", "AvreMagEVCmJE5rEnUXQ9RDWEgZ9cEej12prY4iNYEjr", "Fq6BuKDifozKJUSzMS9i3d7WnSRVPrNQNy4j7i93U3HP", "ENrZXnf9qiRrc5RuoUUBepac7cdqUQo7dGuTUZvmyert", "9b8cirSAK3VWJRNsz2Pm6npJrFfPtGW6XnrGuzgSZP1e", "asdEJnE7osjgnSyQkSZJ3e5YezbmXuDQPiyeyiBxoUm", "DT3SCUT4oN4FkL4iwfY6mNY5qdEekm5MepGQAYAhqjbi", "Fr7cUrU6T66HaLdxJTyXaYY6rdmTvTR5AcoKvo43e8rH", "Fvsp4gb3Qc4Eea7ZpmF7sEgmigfmPCdJ231xeYCRFm4Q", "FihSmFU6J73hxcM7H8XAAEMdBN5V5UCqszEAtyYeF9Ee", "GGQU74M6ikrn8Cj7qywpmj6qdx2nKJLXGb34MbtPChoh", "BSZmKjRxe1SD3QXirFzx4tZvhsKoBTfYNgFJP1XWMCeR", "AVb8sSTQj2ri5WziJu4YhA3N7sMFExfpCsGY1pAvNmPA", "5AtujK4anv4GrEwDUm4WAsHU9zwJsmannQVpDmcY2GQR", "56DgNYguY2YKfdf51jHaX2x7Vay86PZzZCbjDaeBXGci", "4KvuAC2gkGLWoArR5mzNZafMwvgaTajufwPbC5gPvRsE", "Am5Uf31mBgbyVuVBT3je97J4mtCSETuVaJcfxu4C5Tkc", "5DuRdWMtLQ51Ld534PsjDbudPGVnYkCwiGKo5EKcoaaL", "Bjt4YopoMVHD4Bd2hoXWqEZdgwAJfriMdfwkruUkhrvj", "B3Y2Si26KhhBtaj4hgbW4tcY7giNkSrKx8YdRuZP1pVd", "12KcMoB1fHApEx7ghEdTGp43GhvanBCNdQtyavfHqwu8", "2RVjUrDtQVWL4j7nyYx8kDhifmhxAsZM7JRRETm4g9xy", "37NijDJwgTiot4z1BTPZ85jErJbq82d11t8yEWfCk9bN", "8PjGAMT1xXFf7hBu4M4N682AP7TZpXRKzecQ58Y7ceSr", "HassqqjwxLHUUn2RFsNUQSnvxFjPCA84zankooXkr1or", "CbKiQYm6WQPLjC93HLQHcwCQEHamvecQeMMz4Epz5z34", "FeLyAWsSPH1fJto3qKLyQnjCL2j8cVW786RrL3DGf9oY", "AEanj5RgWW67NaRjwHzeETbtvhYwMEkHiW6AceEsq3M4", "EytYPJ4rDdT6dWAdQRw82iqZgezmpKkHQgxCrZXuUqE6", "DF37jxA24hBtyJSkDKZm4ZJL2dBeesEkDwsfjfq7G1a4", "J9FZc62NWM3nxCJm13pzfuwBpAEsPhEYMGsu9QVTP6bi", "EZRHhpvAP4zEX1wZtTQcf6NP4FLWjs9c6tMRBqfrXgFD", "9kV6wT4Kp1AjGYAVRJz2ymK75ytXFK5LKR4X4jFM5Sxh", "GJz3pWa4T3BGxM5dPzCBg1GuTFZ12PuQH9R4eBAqBUKb", "2qg8HoEQZPtphvM1pGcqXp2J8wYS3Y8cRYbnZLcMsABX", "6LdTmuybQ8Lop95MVvfgE7GVyetiQza4rxKam6mni4j2", "Gfv1dbLmGr3r6pmiYtHE2gr3xBC4uwzCgEDJ6Eh1PfHr", "5EotS4QjGN2omyzQer2fPRDZf3xttEtrgrXv53Znfsa", "4iD3dywg6u4geKzvugBn7x1ANDvFjSSLpK4upZKCSLjx", "85UTLhndd3WZrhQvtQfUy4qfnbVggLJmQPB3SqgYjMoL", "FYea5KXFXYbjwoUna6wuJBSqb48UGXmnmGZZBfeVBhfh", "AqNF7zHuccwDzvdUrXrwJwdw6zr3eTSePfzVZpnZLVS8", "D43cyXCLnjGL32wCJUKnU15TpqB7c2ysyurLLxLP8R3F", "Fb69S6d19aKeMn8yJxGfGx1a4nKJsDbdWPmcZgzosjYZ", "ErJHr1hzjKqbsuJjBUrdyetA4syz3oet5cHtZETqHG7u", "CGbid9z39zpTk2vA6yDaHps8s5MVs3ohCzzSdWxkhbAB", "GGjGbs5SePRx4hTmY6ap6jai84m2bPRJ2gYyF4DQ7i8Q", "51QSAAwvEpm11fv92F5xGkR8AW4LyMYdcSP29qgVAaAB", "HM9YXp3qiC6uTRtqmNEpFvLYUw5NUCPr9pCGNxKopC3h", "Gb7YMDtggEYDH8kbgHEn4MFywgnWmEanXrqTr6xEDDxG", "7AZKQo6tRkYGQFBXypN8FHf6EvtZ6QfuUXbhY2E6FoQN", "8gQRste2RKnrhZ7QqZy7hgufkVK6RSAyYjMLhenqbyuh", "Ch21N1nLXFh9DEPGvu5VxWyfnS6C2n8CzuuL2f28bZh5", "AB8jCrJyDHYf25r9aWohdpFFmpGDt6zLNpZc8oJXEPjc", "EU5ZnN5yYzcagJDYH9aMWyKwvcvFg4HLM93Hmidb1ZRC", "8QnyWJaURKqgu8nMxxLYFmAz3APPG4rK7CXFyoQAPgUx", "EcAmT4aE7BP1ooo8ioKbtFcvmvXJ74uJQQ8yi2kkxfJt", "G4m1mm9kMopWxFhkfiUhstXLwyERVZbvvr8F2vPEMr3J", "8NuKR6DAfpaZ14tf2ZaJTXPN924vDbMGCT6rDs8i3bpx", "6LP3CwLwA7StkyMQ9NpKUqLS9ipMmUjPrKhQ8V9w1BoH", "Can9M6kvKYzqm2JcQBckQkUyZhFo3HH1yiiwSE55d4Qg", "BooXF6Jz9kyo3ctMF8Htupx8KVXRjaZioTBEnheUD51a", "ELN2aYLwG92CRCC6XGcWMMb1qtfnQX5deJ8d1cH6V1Zu", "65hXvGkKm2khWRkmkzTzoykM7opSeqBvvpKQ2Hq2HzmX", "2XfxApw3xF1oNYUjHwGGbsMPK6iCc2fvgv1fnCCpUEqT", "7vmiaZEsgaFND5ZcJVdMbwZVi5KJ46zowtY116wJXz8Y", "79dQhwSEfiRh5wpZQrq73JCdt9ghN4kVzJpBWc43v9Zx", "EcRRhPUumt58NFLQnvDCpzFyR8BbkdQApZWR7vJ9BdPo", "661trVCzDWp114gy4PEK4etbjb3u3RNaP4aENa5uN8Vp", "FzXuggFPQVQriHZ6bRU4YiFKj7UTaW1f9uvWWrUZBbnp", "B9YnWRNgboQVqKY3bdJrq4GzmqoB2XdpwReRKQhuNXMP", "33dWwj33J3NUzoTmkMAUq1VdXZL89qezxkdaHdN88vK2", "HRwnH4XVsNKCVX8YhXHv29UXJpkjzFWahWv44PrPR9vK", "Be7uHQhSDZnA4j1E3sfaxtaqjUdoySSJ3kckWEStxPDR", "4nMe3XYXvy3kqnPwLMDZFHfpFu6KHoFwub7oCf3GJEzn", "7FBDhBaVBpNcHiYabHSrB6SZKKE2KH64h9JfSpD9j4NL", "DHoRYvCnFfL53zpq6ZbdHj9wdbtYpK4ip9ieFkk1TyLw", "7qk9vwaHduuQrHkBN9Rtr8ZJwDuozXCtcfGuobJDNHNc", "6H4TkDcHEWkyM2LVNkHdmBsZym4b7Hf5SYfq4HRMbtHR", "F4QCmk3NV3kxYtNwY3d2cSggcYh38uk7fBU4b6rpJ8aZ", "kEhnAVXqmB9qRjZL9WR6pPkcG7EsccnkyUFq8cFRAAX", "FVgEg4dtZ75b5DCETPVQjkhkNdCFKKrGriV5YgtgUAsy", "CzoPGKj2k8WAgijGYfpYxiUJi9cfwcWHiYCit7KPaV5c", "Hbb1Lc7EB6uadY67x1pFaX8inL2y8ABk9rGzX4QLTBkX", "BWUt6TkznDa5kVsXvv6bmGMerjsuBdvF4TEb1Vnuym22",
}

func testGetAccountTime() {
	accounts := make([]solana.PublicKey, 0, len(accountStrSlice))
	for _, accountStr := range accountStrSlice {
		accounts = append(accounts, solana.MustPublicKeyFromBase58(accountStr))
	}
	s := time.Now()
	rpcClient := rpc.New("https://free.rpcpool.com")
	keyedAccounts, err := rpchelp.GetMultipleAccountsWithOpts(rpcClient, context.Background(), accounts, &rpc.GetMultipleAccountsOpts{
		Encoding:   solana.EncodingBase64Zstd,
		Commitment: rpc.CommitmentProcessed,
	})
	fmt.Println(time.Now().Sub(s))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(keyedAccounts))
}
