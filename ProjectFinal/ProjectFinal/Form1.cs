using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace ProjectFinal
{
    public partial class FrmLogin : Form
    {
        public FrmLogin()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            FrmLogPemilik frmLogPemilik = new FrmLogPemilik(); 
            frmLogPemilik.Show();
            this.Hide();
        }

        private void btnPegawai_Click(object sender, EventArgs e)
        {
            FrmLogPegawai frmLogPegawai = new FrmLogPegawai();
            frmLogPegawai.Show();
            this.Hide();
        }
    }
}
